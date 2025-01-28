package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Step 1: Load the airport lookup CSV into a map
	airportData := loadAirportCodes("airport-lookup.csv")

	// Hardcoded input text (replace this with your desired input)
	input := `IATA airports: #WMI, #CEG, #HKI, #SIJ
ICAO airports: ##AYMN, ##UBTT, ##HKII, ##BIHN

            correct, correct, doesnt exist, malformed
IATA cities: *#WMI, *#CEG, *#HKI, *#SIJ
ICAO cities: *##AYMN, *##UBTT, *##HKII, *##BIHN

INCORRECT: ##WMI, #UBTT, *##AYM, *#UBTT, CEG, AYMN, ###WMI, ###AYMN

CORRECT AIRPORTS: Your flight form #ELU to ##FARS is going to be delayed by 1 hour
INCORRECT AIRPORTS: Your flight form ##ELU to #FARS is going to be delayed by 1 hour

CORRECT CITIES: There're 3628km between *#ASE and *##FBMN 
INCORRECT CITIES: There're 3628km between *##ASE and *#FBMN `

	// Step 4: Regex patterns for different codes
	icaoRegex := regexp.MustCompile(`##(\w{4})`)
	cityIcaoRegex := regexp.MustCompile(`\*(##\w{4})`)
	iataRegex := regexp.MustCompile(`#(\w{3})`)
	cityIataRegex := regexp.MustCompile(`\*(#\w{3})`)

	// Process each line of the input text
	for _, line := range strings.Split(input, "\n") {
		processedLine := line

		// Find and replace *##ICAO codes (*##EGLL) with the city name
		processedLine = cityIcaoRegex.ReplaceAllStringFunc(processedLine, func(match string) string {
			code := match[3:]
			if name, found := airportData[code]; found {
				city := getCityFromAirport(name)
				if city != "" {
					return city
				}
			}
			return match
		})

		// Find and replace *#IATA codes (*#LAX) with the city name
		processedLine = cityIataRegex.ReplaceAllStringFunc(processedLine, func(match string) string {
			code := match[2:]
			if name, found := airportData[code]; found {
				city := getCityFromAirport(name)
				if city != "" {
					return city
				}
			}
			return match
		})

		// Find and replace ICAO codes (##EGLL)
		processedLine = icaoRegex.ReplaceAllStringFunc(processedLine, func(match string) string {
			code := match[2:]
			if name, found := airportData[code]; found {
				return name
			}
			return match
		})

		// Find and replace IATA codes (#LAX)
		processedLine = iataRegex.ReplaceAllStringFunc(processedLine, func(match string) string {
			if len(match) == 4 {
				code := match[1:]
				if name, found := airportData[code]; found {
					return name
				}
			}
			return match
		})

		// Print to terminal instead of writing to file
		fmt.Println(processedLine)
	}
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
