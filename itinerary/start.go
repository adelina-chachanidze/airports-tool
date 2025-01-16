package itinerary

import (
	"fmt"
	"os"
	"bufio"
)

func Starting() {
	// Step 1: Open the file
    file, err := os.Open("itinerary/testing.txt") // Open the file named "testing", and showing the path to it
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	 // Step 2: Read the file line by line and print each line
	 scanner := bufio.NewScanner(file)
    
	 //Loop through the file line by line
	 for scanner.Scan() {
		 fmt.Println(scanner.Text()) // Print the current line
	 }
}