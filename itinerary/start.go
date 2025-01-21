package itinerary

import (
	"bufio"
	"fmt"
	"os"
)

func Starting() {
	// Create output.txt at startup
	outputFile, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	inputLines := inputCheck()
	airpotsCheck()
	outputCreate(inputLines)
}

func inputCheck() []string {
	// Open the file named "input.txt", we need to write a path to it
	// file -- the file handle if successful
	// err -- any error that occurred during the operation (or nil if successful)
	file, err := os.Open("itinerary/input.txt")

	//Error handling
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input file not found")
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

	return lines
}

func airpotsCheck() {
	file, err := os.Open("itinerary/airport-lookup.csv")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Airport lookup file not found")
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}

	// for testing, DELETE LATER
	fmt.Println("Airport lookup file found")

	// bufio is connected to the file that was opened by os.Open()
	airportsScanner := bufio.NewScanner(file)

	// Scanner returns true if the next token is available, false otherwise
	if !airportsScanner.Scan() {
		fmt.Println("Airport lookup is empty")
		return
	}

	defer file.Close()
}

func outputCreate(lines []string) {
	file, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	writer.Flush()
	fmt.Println("Output file created successfully")

	cities()
}
