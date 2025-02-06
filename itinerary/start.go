package itinerary

import (
	"bufio"
	"fmt"
	"os"
)

func Starting() {
	// Check for -h flag
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	outputFile, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	inputCheck()
	airportsCheck()

	airportCodes()
}

func inputCheck() {
	file, err := os.Open("itinerary/input.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Input file not found")
		} else {
			fmt.Println("Error opening input file:", err)
		}
		return
	}
	defer file.Close()

	inputScanner := bufio.NewScanner(file)

	// Check if the file is empty
	if !inputScanner.Scan() {
		fmt.Println("The input file is empty")
		return 
	}

	fmt.Println("Input file found successfully")
	return 
}

func airportsCheck() {
	file, err := os.Open("itinerary/airport-lookup.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Airport lookup file not found")
		} else {
			fmt.Println("Error opening airport lookup file:", err)
		}
		return 
	}
	defer file.Close()

	airportsScanner := bufio.NewScanner(file)

	// Skip header row (first line)
	if !airportsScanner.Scan() {
		fmt.Println("Airport lookup file is empty")
		return 
	}

	fmt.Println("Airport lookup file found successfully")
	return 
}
