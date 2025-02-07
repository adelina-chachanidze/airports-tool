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

	// Convert content to string for easier manipulation
	text := string(content)

	// Replace line-break characters with newline
	text = strings.ReplaceAll(text, "\v", "\n")
	text = strings.ReplaceAll(text, "\f", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	// Split into lines for processing
	lines := strings.Split(text, "\n")

	// Process each line and handle consecutive blank lines
	var processedLines []string
	lastLineWasBlank := false

	for _, line := range lines {
		// Trim extra spaces within the line
		line = strings.Join(strings.Fields(line), " ")

		isBlankLine := len(strings.TrimSpace(line)) == 0

		// Skip if we would create consecutive blank lines
		if isBlankLine && lastLineWasBlank {
			continue
		}

		processedLines = append(processedLines, line)
		lastLineWasBlank = isBlankLine
	}

	// Join lines back together and write to file
	output := strings.Join(processedLines, "\n")
	err = os.WriteFile("itinerary/output.txt", []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
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
		line := scanner.Text()
		lineNumber++

		// Check for '#' or non-ASCII characters
		if strings.Contains(line, "#") || containsNonASCII(line) {
			errorLines = append(errorLines, lineNumber)
		}
	}

	if len(errorLines) > 0 {
		numbers := make([]string, len(errorLines))
		for i, line := range errorLines {
			numbers[i] = fmt.Sprintf("%d", line)
		}
		fmt.Printf("\033[33mPossible errors were detected on line(s) %s in the output file. Please check if formatting is correct in the input file.\033[0m\n",
			strings.Join(numbers, ","))
	}
}

func containsNonASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return true
		}
	}
	return false
}
