package itinerary

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func outputContents() {
	// Step 1: Load the airport lookup CSV into a map
	airportData := loadAirportData("itinerary/airport-lookup.csv")

	// Step 2: Open the input file
	file, _ := os.Open("itinerary/input.txt")
	defer file.Close()

	// Step 3: Create the output file
	outputFile, _ := os.Create("itinerary/output.txt")
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(file)

	// Step 4: Regex patterns for IATA and ICAO codes
	iataRegex := regexp.MustCompile(`#(\w{3})`)    // Matches IATA codes (#LAX)
	icaoRegex := regexp.MustCompile(`##(\w{4})`)  // Matches ICAO codes (##EGLL)

	// Step 5: Process each line of the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Find and replace IATA codes
		line = iataRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[1:] // Extract the code (e.g., "LAX" from "#LAX")
			if name, found := airportData[code]; found {
				return name // Replace with the airport name
			}
			return match // Leave unchanged if not found
		})

		// Find and replace ICAO codes
		line = icaoRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[2:] // Extract the code (e.g., "EGLL" from "##EGLL")
			if name, found := airportData[code]; found {
				return name // Replace with the airport name
			}
			return match // Leave unchanged if not found
		})

		// Write the updated line to the output file
		writer.WriteString(line + "\n")
	}

	writer.Flush()
	fmt.Println("Output file created with airport names")
}

// Helper function to load airport data into a map
func loadAirportData(filepath string) map[string]string {
	file, _ := os.Open(filepath)
	defer file.Close()

	reader := csv.NewReader(file)
	airportMap := make(map[string]string)

	// Read the header row (skip it)
	_, _ = reader.Read()

	// Read each row
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Ensure all columns are present and valid
		if len(record) < 5 || strings.TrimSpace(record[3]) == "" || strings.TrimSpace(record[4]) == "" {
			continue // Skip malformed rows
		}

		icaoCode := strings.TrimSpace(record[3]) // Column 4: ICAO code
		iataCode := strings.TrimSpace(record[4]) // Column 5: IATA code
		name := strings.TrimSpace(record[0])     // Column 1: Airport name

		// Add to the map
		airportMap[icaoCode] = name
		airportMap[iataCode] = name
	}

	return airportMap
}
