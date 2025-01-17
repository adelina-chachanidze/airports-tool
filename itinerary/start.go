package itinerary

import (
	"bufio"
	"fmt"
	"os"
)

func Starting() {
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

	defer file.Close()

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
	file, err := os.Open("itineraryairport-lookup.csv")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Airport lookup file not found")
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Println("Airport lookup file found")

	airportsScanner := bufio.NewScanner(file)

	if !airportsScanner.Scan() {
		fmt.Println("Airport lookup is empty")
		return
	}

	defer file.Close()
}

func outputCreate() {
	os.Create("output.txt")
}
