package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractWords(s string) []string {
	parts := strings.Split(s, "#")
	var words []string
	for _, word := range parts {
		if len(word) >= 2 {
			words = append(words, word)
		}
	}
	return words
}

func main() {
	rows, err1 := strconv.Atoi(os.Args[1])
	cols, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Please provide valid numbers for rows and columns.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	matrix := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			fmt.Println("Failed to read input for row", i+1)
			return
		}

		line := scanner.Text()
		if len(line) != cols {
			fmt.Printf("Row %d should have %d characters but got %d\n", i+1, cols, len(line))
			return
		}

		matrix[i] = make([]rune, cols)
		for j, ch := range line {
			matrix[i][j] = ch
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var words []string

	// Extract words from rows
	for _, row := range matrix {
		rowWord := string(row)
		words = append(words, extractWords(rowWord)...)
	}

	// Extract words from columns
	for j := 0; j < cols; j++ {
		colWord := ""
		for i := 0; i < rows; i++ {
			colWord += string(matrix[i][j])
		}
		words = append(words, extractWords(colWord)...)
	}

	// Print the words
	// for _, word := range words {
	// 	fmt.Println(word)
	// }

	// // Sort the words lexicographically
	// sort.Strings(words)

	// Check if we have any words and print the smallest one
	if len(words) > 0 {
		fmt.Println("The lexicographically smallest word is:", words[0])
	} else {
		fmt.Println("No valid words found.")
	}
}
