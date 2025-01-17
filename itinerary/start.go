package itinerary

import (
	"bufio"
	"fmt"
	"os"
)

func Starting() {
	inputCheck()
	airpotsCheck()
	outputCreate()
}

func inputCheck() {
	// Open the file named "testing", we need to write a path to it
	// file -- the file handle if successful
	// err -- any error that occurred during the operation (or nil if successful)
	file, err := os.Open("itinerary/testing.txt")

	//Error handling
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input file not found")
		} else {
			fmt.Println("Error opening input file:", err)
		}
		return
	}

	//Read the file line by line and print each line
	inputScanner := bufio.NewScanner(file)

	//Check if the file is empty
	if !inputScanner.Scan() {
		fmt.Println("The input file is empty")
		return
	}

	//Loop through the file line by line
	for inputScanner.Scan() {
		fmt.Println(inputScanner.Text())
		// Returns the current line of text that the scanner has read as a string
		// The text is returned as a string
		// Part of the bufio package
	}
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

func outputCreate() {
	file, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Output file created successfully")
}
