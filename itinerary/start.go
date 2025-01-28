package itinerary

import (
	"bufio"
	"fmt"
	"os"
)

func Starting() {
	// Check if output file already exists and remove error handling since we'll handle it differently
	if _, err := os.Stat("itinerary/output.txt"); err == nil {
		os.Remove("itinerary/output.txt")
	}

	// Check input file first
	lines := inputCheck()
	if lines == nil {
		return // Don't create output file if input check fails
	}

	// Check airport lookup file
	if !airpotsCheck() {
		return // Don't create output file if airport check fails
	}

	// Only create output file if all checks pass
	outputFile, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write to the output file
	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}
	writer.Flush() // Don't forget to flush the buffer
}

func inputCheck() []string {
	// Open the file named "input.txt", we need to write a path to it
	// file -- the file handle if successful
	// err -- any error that occurred during the operation (or nil if successful)
	file, err := os.Open("itinerary/input.txt")

	//Error handling
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input not found")
		} else {
			fmt.Println("Error opening input file:", err)
		}
		return nil
	}
	defer file.Close()

	var lines []string
	inputScanner := bufio.NewScanner(file)

	//Check if the file is empty
	if !inputScanner.Scan() {
		fmt.Println("The input file is empty")
		return nil
	}

	lines = append(lines, inputScanner.Text())

	//Loop through the file line by line
	for inputScanner.Scan() {
		lines = append(lines, inputScanner.Text())
	}

	fmt.Println("Input file found")
	return lines
}

func airpotsCheck() bool {
	file, err := os.Open("itinerary/airport-lookup.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Airport lookup not found")
			return false
		}
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// for testing, DELETE LATER
	fmt.Println("Airport lookup file found")

	// bufio is connected to the file that was opened by os.Open()
	airportsScanner := bufio.NewScanner(file)

	// Scanner returns true if the next token is available, false otherwise
	if !airportsScanner.Scan() {
		fmt.Println("Airport lookup is empty")
		return false
	}

	// Check if header row has correct columns
	expectedHeaders := []string{"name", "iso_country", "municipality", "icao_code", "iata_code", "coordinates"}
	headers := splitCSVLine(airportsScanner.Text())

	if len(headers) != len(expectedHeaders) {
		fmt.Println("Airport lookup malformed")
		return false
	}

	for i, header := range headers {
		if header == "" || header != expectedHeaders[i] {
			fmt.Println("Airport lookup malformed")
			return false
		}
	}

	airportCodes()
	return true
}

// Helper function to split CSV line considering possible commas within quoted fields
func splitCSVLine(line string) []string {
	var result []string
	var current string
	inQuotes := false

	for _, char := range line {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ',':
			if !inQuotes {
				result = append(result, current)
				current = ""
				continue
			}
		}
		current += string(char)
	}
	result = append(result, current)
	return result
}

/*func outputCreate() {
	file, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Output file created")
	processAirportCodes()
}*/
