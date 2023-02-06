package services

import (
	"fmt"
	"os"
)

func CheckExistDir(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist, creating...\n", dirPath)
		// Create new directory with 0755 permissions
		if err := os.Mkdir(dirPath, 0755); err != nil {
			fmt.Printf("Error creating directory: %s\n", err)
			return false
		}
	} else {
		fmt.Printf("Directory %s already exists\n", dirPath)
	}
	return true
}
