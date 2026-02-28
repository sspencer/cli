package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var reader io.Reader

	// Check if a filename argument is provided
	if len(os.Args) > 1 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		// Read from stdin
		reader = os.Stdin
	}

	// Read all input
	scanner := bufio.NewScanner(reader)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text() + " ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Extract all positive integers using regex
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text.String(), -1)

	// Convert to integers and filter valid positive integers
	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil && num > 0 {
			numbers = append(numbers, num)
		}
	}

	// Output as comma-separated values without newline
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
}
