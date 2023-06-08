package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get the default temporary directory path
	tempPath := os.TempDir() // different depending on the OS
	fmt.Println("The default temp dir path is ", tempPath)

	// create a temporary file use TempFile
	tempContent := []byte("This is some temp content for the file")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt") // * is replaced with random string by TempFile
	if err != nil {
		panic(err)
	}
	// Write the tempContent to the tempFile
	if _, err := tmpFile.Write(tempContent); err != nil {
		panic(err)
	}
	// Read and print the tempFile content
	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Printf("%s\n", data)

	// Remove the temp file when finished
	fmt.Println("Temp file is named:", tmpFile.Name())
	defer os.Remove(tmpFile.Name())

	// create a temporary directory with ioutil's version of TempDir
	tmpDir, err := ioutil.TempDir("", "tempdir*") // * is replaced with random string by TempDir
	if err != nil {
		panic(err)
	}
	fmt.Println("The temp dir is named:", tmpDir)
	defer os.RemoveAll(tmpDir)
}
