package itinerary

import (
	"fmt"
	"os"
)

func cities() {
	data, err := os.ReadFile("itinerary/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = os.WriteFile("itinerary/output.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File copied successfully")
}
