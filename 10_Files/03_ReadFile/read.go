package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getFileLength(filepath string) int64 {
	f, err := os.Stat(filepath)
	handleErr(err)
	return f.Size()
}

func main() {
	// Read Entire file (not always enough memory for large files)
	content, err := ioutil.ReadFile("sampletext.txt")
	handleErr(err)
	// ReadFile reads the data as bytes, we have to convert to a string
	fmt.Println(string(content))

	f, _ := os.Open("sampletext.txt")
	defer f.Close()

	// Read file in pieces with buffer
	const BuffSize = 20
	b1 := make([]byte, BuffSize)

	for {
		n, err := f.Read(b1)

		if err != nil {
			if err != io.EOF {
				handleErr(err)
			}
			break // break when EOF
		}

		fmt.Println("Bytes read:", n)
		fmt.Println("Content:", string(b1[:n]))
	}

	// Get the length of a file
	l := getFileLength("sampletext.txt")
	fmt.Println("File length is:", l)

	// Use ReadAt to read from a specific index
	b2 := make([]byte, 10) // 10 is both length and capacity
	// read starting from the offset l-int64(len(b2)), that is from end of file backward by 10 bytes
	_, err3 := f.ReadAt(b2, l-int64(len(b2)))
	handleErr(err3)
	fmt.Println(string(b2))
}
