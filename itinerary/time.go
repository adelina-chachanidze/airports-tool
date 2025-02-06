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
	input, err := os.ReadFile("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error reading output file:", err)
		return
	}
	lines := strings.Split(string(input), "\n")

	var processedLines []string
	seenTimeSection := false

	for _, line := range lines {
		if strings.HasPrefix(line, "1. D(") || strings.HasPrefix(line, "1. T") {
			seenTimeSection = true
			break
		}
		if !seenTimeSection {
			processedLines = append(processedLines, line)
		}
	}

	timeLines := getTimeEntries(lines)
	processedLines = append(processedLines, timeLines...)

	outputFile, err := os.Create("itinerary/output.txt")
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range processedLines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
}

func getTimeEntries(lines []string) []string {
	var timeEntries []string
	for _, line := range lines {
		if strings.HasPrefix(line, "1. ") ||
			strings.HasPrefix(line, "2. ") ||
			strings.HasPrefix(line, "3. ") ||
			strings.HasPrefix(line, "4. ") ||
			strings.HasPrefix(line, "5. ") ||
			strings.HasPrefix(line, "6. ") ||
			strings.HasPrefix(line, "7. ") ||
			strings.HasPrefix(line, "8. ") ||
			strings.HasPrefix(line, "9. ") {

			line = processTimeFormats(line)
			timeEntries = append(timeEntries, line)
		}
	}
	return timeEntries
}

func processTimeFormats(line string) string {
	time12Regex := regexp.MustCompile(`T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00)\)`)
	zulu12Regex := regexp.MustCompile(`T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}Z)\)`)
	time24Regex := regexp.MustCompile(`T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}[+-]\d{2}:00)\)`)
	zulu24Regex := regexp.MustCompile(`T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}Z)\)`)

	line = time12Regex.ReplaceAllStringFunc(line, func(match string) string {
		if timeStr := extractAndFormat12HourTime(match[4 : len(match)-1]); timeStr != "" {
			return timeStr
		}
		return match
	})

	line = zulu12Regex.ReplaceAllStringFunc(line, func(match string) string {
		if timeStr := extractAndFormat12HourTime(match[4 : len(match)-1]); timeStr != "" {
			return timeStr
		}
		return match
	})

	line = time24Regex.ReplaceAllStringFunc(line, func(match string) string {
		if timeStr := extractAndFormat24HourTime(match[4 : len(match)-1]); timeStr != "" {
			return timeStr
		}
		return match
	})

	line = zulu24Regex.ReplaceAllStringFunc(line, func(match string) string {
		if timeStr := extractAndFormat24HourTime(match[4 : len(match)-1]); timeStr != "" {
			return timeStr
		}
		return match
	})

	return line
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
