package itinerary

import (
	"fmt"
	"os"
	//"regexp"
	//"strings"
	"time"
	"bufio"
)


func formatDate() error {
	// Open the input file
	inputFile, err := os.Open("itinerary/input.txt")
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	// Open the existing output file in write mode
	outputFile, err := os.OpenFile("itinerary/output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("error opening output file: %w", err)
	}
	defer outputFile.Close()

	// Use a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the date string using the format D(yyyy-mm-dd)
		layout := "D(2006-01-02)"
		t, err := time.Parse(layout, line)
		if err != nil {
			return fmt.Errorf("error formatting date '%s': %w", line, err)
		}

		// Format the date into "dd MMM yyyy"
		formattedDate := t.Format("02 Jan 2006")

		// Write the formatted date to the output file
		_, err = outputFile.WriteString(formattedDate + "\n")
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	return nil
}

/*func formatDate() (string, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("error reading input file: %v", err)
	}

	input := strings.TrimSpace(string(content))
	// Parse the date string using the format D(yyyy-mm-dd)
	layout := "D(2006-01-02)"
	t, err := time.Parse(layout, input)
	if err != nil {
		return "", err
	}

	// Format the date into "dd MMM yyyy"
	return t.Format("02 Jan 2006"), nil


}*/

/*func formatDateTime() (string, error) {
	// Read from input.txt
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("error reading input file: %v", err)
	}

	input := strings.TrimSpace(string(content))

	// Regular expressions for validation
	datePattern := regexp.MustCompile(`^D\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?:Z|[+-]\d{2}:\d{2}))\)$`)
	time12Pattern := regexp.MustCompile(`^T12\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?:Z|[+-]\d{2}:\d{2}))\)$`)
	time24Pattern := regexp.MustCompile(`^T24\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?:Z|[+-]\d{2}:\d{2}))\)$`)

	var (
		formatType string
		dtString   string
	)

	// Check which pattern matches and extract the datetime string
	switch {
	case datePattern.MatchString(input):
		formatType = "date"
		dtString = datePattern.FindStringSubmatch(input)[1]
	case time12Pattern.MatchString(input):
		formatType = "time12"
		dtString = time12Pattern.FindStringSubmatch(input)[1]
	case time24Pattern.MatchString(input):
		formatType = "time24"
		dtString = time24Pattern.FindStringSubmatch(input)[1]
	default:
		return input, nil
	}

	// Handle 'Z' timezone
	if strings.HasSuffix(dtString, "Z") {
		dtString = strings.TrimSuffix(dtString, "Z") + "+00:00"
	}

	// Parse the datetime
	dt, err := time.Parse("2006-01-02T15:04-07:00", dtString)
	if err != nil {
		return input, nil
	}

	// Extract timezone offset
	offset := dt.Format("-07:00")
	if offset == "+00:00" {
		offset = "Z"
	}

	// Format according to type
	switch formatType {
	case "date":
		return dt.Format("02 Jan 2006"), nil
	case "time12":
		return fmt.Sprintf("%s (%s)", dt.Format("03:04PM"), offset), nil
	case "time24":
		return fmt.Sprintf("%s (%s)", dt.Format("15:04"), offset), nil
	default:
		return input, nil
	}
}*/
