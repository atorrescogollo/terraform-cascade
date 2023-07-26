package utils

import (
	"fmt"
	"log"
	"strings"
)

// WaitForConfirmation waits for user confirmation (y/n)
func WaitForConfirmation(message string) bool {
	var response string

	fmt.Println(message + ":")
	for {
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}

		switch strings.ToLower(response) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Error: Please type (y)es or (n)o and then press enter:")
		}
	}
}
