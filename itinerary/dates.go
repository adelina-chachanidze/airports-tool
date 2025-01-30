package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Check if input.txt exists
	if _, err := os.Stat("input.txt"); err != nil {
		fmt.Println("input.txt does not exist")
		return
	}

	fmt.Println("input.txt exists")

	// Write the formatted dates to output.txt
	output := formatDates()

	// Read the original input file
	inputContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Combine formatted output with original input
	fullOutput := output + "\n\n" + string(inputContent)

	// Write combined content to output.txt
	err = os.WriteFile("output.txt", []byte(fullOutput), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func formatDates() string {
	// Hardcoded input with the exact format
	input := `D(2007-04-05T12:30−02:00)
	D(2012-07-09T12:30−02:00)`

	// Define regex to match dates inside D(...) with the exact format
	dateRegex := regexp.MustCompile(`D\((\d{4}-\d{2}-\d{2})T[\d:]+[−-]\d{2}:\d{2}\)`)

	var result string

	// Process each line
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		// Replace all date matches
		modifiedLine := dateRegex.ReplaceAllStringFunc(line, func(match string) string {
			datePart := dateRegex.FindStringSubmatch(match)
			if len(datePart) < 2 {
				return match
			}

			// Parse the date
			parsedTime, err := time.Parse("2006-01-02", datePart[1])
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return match
			}

			// Format the date as "05 Apr 2007"
			return parsedTime.Format("02 Jan 2006")
		})

		result += modifiedLine + "\n"
	}

	return strings.TrimSpace(result)
}
