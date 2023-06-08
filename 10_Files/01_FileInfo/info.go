package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	// golang.org/pkg/os/#FileInfo
	// os.State returns FileInfo struct, which contains fields Name(), Size(), Mode(), ModTime(), IsDir(), Sys()
	if _, err := os.Stat(filePath); err != nil {
		// os.IsNotExist returns a boolean value indicating whether the error is known to report that a file or directory does not exist.
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func main() {
	// Use the Stat function to get file stats
	// stats is a FileInfo struct
	stats, err := os.Stat("sampletext.txt")
	if err != nil {
		panic(err)
	}

	// Check if a file exists
	exists := checkFileExists("sampletext.txt")
	fmt.Println("File exists check:", exists) // File exists check: true

	// Get the file's modification time
	fmt.Println("Modification time:", stats.ModTime())
	// Modification time: 2023-06-07 11:40:02.9572065 +0200 CEST

	fmt.Println("File mode:", stats.Mode()) // File mode: -rw-rw-rw-
	fmode := stats.Mode()
	if fmode.IsRegular() {
		fmt.Println("This is a regular file")
	}
	// This is a regular file

	// Get the file size
	fmt.Println("The length of the file is:", stats.Size())
	// The length of the file is: 0

	// Is this a directory?
	fmt.Println("Is directory:", stats.IsDir())
	// Is directory: false
}
