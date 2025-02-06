package itinerary

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func airportCodes() {
	// Step 1: Load the airport lookup CSV into a map
	airportData := loadAirportCodes("itinerary/airport-lookup.csv")

	// Step 2: Open the input file
	file, _ := os.Open("itinerary/input.txt")
	defer file.Close()

	// Step 3: Open the output file in append mode
	outputFile, _ := os.OpenFile("itinerary/output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(file)

	// Step 4: Regex patterns for different codes
	icaoRegex := regexp.MustCompile(`\#\#(\w{4})\b`)
	cityIcaoRegex := regexp.MustCompile(`\*(##\w{4})`)
	iataRegex := regexp.MustCompile(`(^|[^#])#(\w{3})\b`)
	cityIataRegex := regexp.MustCompile(`\s\*#(\w{3})\b`)

	// Step 5: Process each line of the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Find and replace *##ICAO codes (*##EGLL) with the city name
		line = cityIcaoRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[3:] // Extract the code (e.g., "EGLL" from "*##EGLL")
			if name, found := airportData[code]; found {
				city := getCityFromAirport(name)
				if city != "" {
					return city // Replace with the city name
				}
			}
			return match // Leave unchanged if not found or invalid
		})

		// Find and replace *#IATA codes (*#LAX) with the city name
		line = cityIataRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[2:] // Extract the code (e.g., "LAX" from "*#LAX")
			if name, found := airportData[code]; found {
				city := getCityFromAirport(name)
				if city != "" {
					return city // Replace with the city name
				}
			}
			return match // Leave unchanged if not found or invalid
		})

		// Find and replace ICAO codes (##EGLL)
		line = icaoRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[2:] // Extract the code (e.g., "EGLL" from "##EGLL")
			if name, found := airportData[code]; found {
				return name // Replace with the airport name
			}
			return match // Leave unchanged if not found
		})

		// Find and replace IATA codes (#LAX)
		line = iataRegex.ReplaceAllStringFunc(line, func(match string) string {
			code := match[2:] // Extract the code (e.g., "LAX" from "#LAX")
			if name, found := airportData[code]; found {
				return name // Replace with the airport name
			}
			return match // Leave unchanged if not found
		})

		// Write the updated line to the output file
		writer.WriteString(line + "\n")
	}

	writer.Flush()
	fmt.Println("Output file created with replacements")

	formatTimes()

	outputFormatting()
}

// Helper function to load airport data into a map
func loadAirportCodes(filepath string) map[string]string {
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

	fmt.Println("Processing csv")
	return airportMap
}

// Helper function to extract the city name from the airport name
func getCityFromAirport(airportName string) string {
	// Assuming the city name is the first part of the airport name, separated by a space
	parts := strings.Split(airportName, " ")
	if len(parts) > 0 {
		return parts[0] // Return the first word as the city name
	}
	return ""
}
