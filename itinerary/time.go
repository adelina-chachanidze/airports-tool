package itinerary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func formatDates() {
	// Open the input and output files
	input, _ := os.ReadFile("itinerary/output.txt")
	lines := strings.Split(string(input), "\n")

	outputFile, _ := os.Create("itinerary/output.txt")
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	fmt.Println("Testing formatDates")

	// Regex patterns for times
	time12Regex := regexp.MustCompile(`T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00|)\)`)
	time24Regex := regexp.MustCompile(`T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00|)\)`)

	for _, line := range lines {
		// Process 12-hour times
		line = time12Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat12HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		// Process 24-hour times
		line = time24Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat24HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		writer.WriteString(line + "\n")
	}

	writer.Flush()
}

func extractAndFormat12HourTime(isoDate string) string {
	if isoDate == "" {
		return ""
	}

	t, err := time.Parse("2006-01-02T15:04-07:00", isoDate)
	if err != nil {
		return ""
	}

	hour := t.Hour()
	ampm := "AM"
	if hour >= 12 {
		ampm = "PM"
		if hour > 12 {
			hour -= 12
		}
	}
	if hour == 0 {
		hour = 12
	}

	_, offset := t.Zone()
	offsetHours := offset / 3600
	offsetStr := fmt.Sprintf("%+03d:00", offsetHours)

	return fmt.Sprintf("%d:%02d%s (%s)", hour, t.Minute(), ampm, offsetStr)
}

func extractAndFormat24HourTime(isoDate string) string {
	if isoDate == "" {
		return ""
	}

	t, err := time.Parse("2006-01-02T15:04-07:00", isoDate)
	if err != nil {
		return ""
	}

	_, offset := t.Zone()
	offsetHours := offset / 3600
	offsetStr := fmt.Sprintf("%+03d:00", offsetHours)

	return fmt.Sprintf("%02d:%02d (%s)", t.Hour(), t.Minute(), offsetStr)
}
