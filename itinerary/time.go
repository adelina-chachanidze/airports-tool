package itinerary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func formatTimes() {
	// Open the input and output files
	input, _ := os.ReadFile("itinerary/output.txt")
	lines := strings.Split(string(input), "\n")

	outputFile, _ := os.Create("itinerary/output.txt")
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	// Regex patterns for times and dates
	time12Regex := regexp.MustCompile(`T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00)\)`)
	zulu12Regex := regexp.MustCompile(`T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}Z)\)`)
	time24Regex := regexp.MustCompile(`T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00)\)`)
	zulu24Regex := regexp.MustCompile(`T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}Z)\)`)

	for _, line := range lines {
		// Process 12-hour offset times
		line = time12Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat12HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		// Process 12-hour Zulu times
		line = zulu12Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat12HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		// Process 24-hour offset times
		line = time24Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat24HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		// Process 24-hour Zulu times
		line = zulu24Regex.ReplaceAllStringFunc(line, func(match string) string {
			if timeStr := extractAndFormat24HourTime(match[4 : len(match)-1]); timeStr != "" {
				return timeStr
			}
			return match
		})

		writer.WriteString(line + "\n")
	}

	formatDates()
	writer.Flush()
}

// Convert ISO date to 12-hour format with AM/PM notation
func extractAndFormat12HourTime(isoDate string) string {
	if isoDate == "" {
		return ""
	}

	var t time.Time
	var err error

	// Parse Zulu format
	if strings.HasSuffix(isoDate, "Z") {
		t, err = time.Parse("2006-01-02T15:04Z", isoDate)
	} else {
		t, err = time.Parse("2006-01-02T15:04-07:00", isoDate)
	}

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

	// Adjust for Zulu time (UTC)
	offsetStr := getOffsetString(isoDate, t)

	return fmt.Sprintf("%02d:%02d%s (%s)", hour, t.Minute(), ampm, offsetStr)
}

// Convert ISO date to 24-hour format
func extractAndFormat24HourTime(isoDate string) string {
	if isoDate == "" {
		return ""
	}

	var t time.Time
	var err error

	// Parse Zulu format
	if strings.HasSuffix(isoDate, "Z") {
		t, err = time.Parse("2006-01-02T15:04Z", isoDate)
	} else {
		t, err = time.Parse("2006-01-02T15:04-07:00", isoDate)
	}

	if err != nil {
		return ""
	}

	// Adjust for Zulu time (UTC)
	offsetStr := getOffsetString(isoDate, t)

	return fmt.Sprintf("%02d:%02d (%s)", t.Hour(), t.Minute(), offsetStr)
}

// Returns offset in ±hh:00 format, ensuring Zulu time is displayed as "+00:00"
func getOffsetString(isoDate string, t time.Time) string {
	if strings.HasSuffix(isoDate, "Z") {
		return "+00:00"
	}
	_, offset := t.Zone()
	offsetHours := offset / 3600
	return fmt.Sprintf("%+03d:00", offsetHours)
}
