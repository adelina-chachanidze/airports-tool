package itinerary

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func airportCodes() {
	airportData := loadAirportCodes("itinerary/airport-lookup.csv")

	// Read the entire output file
	input, err := os.ReadFile("itinerary/output.txt")
	if err != nil {
		return
	}
	content := string(input)

	// Define regex patterns
	icaoRegex := regexp.MustCompile(`\#\#(\w{4})\b`)
	cityIcaoRegex := regexp.MustCompile(`\*(##\w{4})`)
	iataRegex := regexp.MustCompile(`(^|[^#])#(\w{3})\b`)
	cityIataRegex := regexp.MustCompile(`\s\*#(\w{3})\b`)

	// Process the content with all regex replacements
	content = cityIcaoRegex.ReplaceAllStringFunc(content, func(match string) string {
		code := match[3:]
		if name, found := airportData[code]; found {
			city := getCityFromAirport(name)
			if city != "" {
				return city
			}
		}
		return match
	})

	content = cityIataRegex.ReplaceAllStringFunc(content, func(match string) string {
		code := match[2:]
		if name, found := airportData[code]; found {
			city := getCityFromAirport(name)
			if city != "" {
				return " " + city
			}
		}
		return match
	})

	content = icaoRegex.ReplaceAllStringFunc(content, func(match string) string {
		code := match[2:]
		if name, found := airportData[code]; found {
			return name
		}
		return match
	})

	content = iataRegex.ReplaceAllStringFunc(content, func(match string) string {
		code := match[2:]
		if name, found := airportData[code]; found {
			return name
		}
		return match
	})

	// Write back to the output file
	os.WriteFile("itinerary/output.txt", []byte(content), 0644)

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
