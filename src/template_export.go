package src

import (
	"os"
)

// Write to files to Output folder
func CreateBundles(filename string) (*os.File, error) {
	// Open the file for writing (create it if it doesn't exist)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
