package itinerary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Starting() {
	// Check for -h flag
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	if inputCheck() && airportsCheck() {
		createOutputFile()
	}
}

func createOutputFile() {
	// Open input file
	inputFile, err := os.Open("itinerary/input.txt")
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Copy contents from input to output
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
		return
	}
	writer.Flush()
	airportCodes()
}

func inputCheck() bool {
	file, err := os.Open("itinerary/input.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input file not found")
		} else {
			fmt.Println("Error opening input file:", err)
		}
		return false
	}
	defer file.Close()

	inputScanner := bufio.NewScanner(file)

	// Check if the file is empty
	if !inputScanner.Scan() {
		fmt.Println("The input file is empty")
		return false
	}

	fmt.Println("Input file found successfully")
	return true

}

func airportsCheck() bool {
	file, err := os.Open("itinerary/airport-lookup.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Airport lookup file not found")
		} else {
			fmt.Println("Error opening airport lookup file:", err)
		}
		return false
	}
	defer file.Close()

	airportsScanner := bufio.NewScanner(file)

	// Skip header row (first line)
	if !airportsScanner.Scan() {
		fmt.Println("Airport lookup file is empty")
		return false
	}

	// Check header row format
	headerRow := airportsScanner.Text()
	expectedHeader := "name,iso_country,municipality,icao_code,iata_code,coordinates"
	if headerRow != expectedHeader {
		fmt.Println("Airport lookup file is malformed: incorrect column order or missing columns")
		return false
	}

	// Check each row for empty cells
	lineNumber := 1 // Start from 1 since we already read the header
	for airportsScanner.Scan() {
		lineNumber++
		line := airportsScanner.Text()

		// Split the line into columns
		columns := strings.Split(line, ",")

		// Check for empty cells
		for i, cell := range columns {
			if strings.TrimSpace(cell) == "" {
				fmt.Printf("Airport lookup file is malformed: empty cell found on line %d, column %d\n", lineNumber, i+1)
				return false
			}
		}
	}

	if err := airportsScanner.Err(); err != nil {
		fmt.Println("Error reading airport lookup file:", err)
		return false
	}

	fmt.Println("Airport lookup file found successfully")
	return true
}
