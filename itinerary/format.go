package itinerary

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func outputFormatting() error {
	// Read the entire output file
	content, err := os.ReadFile("itinerary/output.txt")
	if err != nil {
		return fmt.Errorf("error reading output file: %w", err)
	}

	// Convert vertical whitespace characters to newlines
	content = bytes.ReplaceAll(content, []byte{'\v'}, []byte{'\n'})
	content = bytes.ReplaceAll(content, []byte{'\f'}, []byte{'\n'})
	content = bytes.ReplaceAll(content, []byte{'\r'}, []byte{'\n'})

	// Create a scanner to process line by line
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var formattedLines []string
	emptyLineCount := 0

	// Process each line
	for scanner.Scan() {
		// Trim trailing whitespace and collapse multiple spaces into single spaces
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), " ")

		if line == "" {
			emptyLineCount++
			if emptyLineCount <= 1 {
				formattedLines = append(formattedLines, line)
			}
		} else {
			emptyLineCount = 0
			formattedLines = append(formattedLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning content: %w", err)
	}

	// Write the formatted content back to the file
	output := strings.Join(formattedLines, "\n")
	err = os.WriteFile("itinerary/output.txt", []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error writing formatted content: %w", err)
	}

	userErrors()

	return nil
}

func userErrors() {
	content, _ := os.ReadFile("itinerary/output.txt")
	scanner := bufio.NewScanner(bytes.NewReader(content))
	lineNumber := 0
	var errorLines []int

	for scanner.Scan() {
		lineNumber++
		if strings.Contains(scanner.Text(), "#") {
			errorLines = append(errorLines, lineNumber)
		}
	}

	if len(errorLines) > 0 {
		numbers := make([]string, len(errorLines))
		for i, line := range errorLines {
			numbers[i] = fmt.Sprintf("%d", line)
		}
		fmt.Printf("\033[33mPossible error on line(s) %s. Please check your input file.\033[0m\n",
			strings.Join(numbers, ","))
	}
}
