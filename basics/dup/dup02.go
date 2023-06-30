/*
This Go program counts the occurrences of each line of text in one or more input files
or from the standard input. It then prints the lines that appear more than once, along
with the number of times they appear.

Algorithm:
 1. Create an empty map called "counts" to store the counts of each line.
 2. Get the command-line arguments (file names) excluding the program name.
 3. If no file names are provided:
    3.1. Count the lines from the standard input using the "countLines" function.
 4. If file names are provided:
    4.1. Iterate over each file name.
    4.1.1. Open the file.
    4.1.2. If there is an error opening the file, print the error and continue to the next file.
    4.1.3. Count the lines from the opened file using the "countLines" function.
    4.1.4. Close the file.
 5. Iterate over the "counts" map.
    5.1. For each line and its corresponding count in the "counts" map:
    5.1.1. Check if the count (number of occurrences) is greater than 1.
    5.1.1.1. If yes, print the number of occurrences and the line.
 6. End of the program.

This program allows counting the occurrences of lines in one or more files or from the standard input. It uses a map to store the counts, and the "countLines" function is responsible for counting the lines from a given source.
*/
package main

import (
	"bufio" // Package for reading input
	"fmt"   // Package for formatting and printing
	"os"    // Package for interacting with the operating system
)

// The main function is the entry point of the program.
func main() {
	// Create an empty map to store the counts of each line.
	counts := make(map[string]int)

	// Get the command-line arguments (file names) excluding the program name.
	files := os.Args[1:]

	// If no file names are provided:
	if len(files) == 0 {
		// Count the lines from the standard input using the "countLines" function.
		countLines(os.Stdin, counts)
	} else {
		// If file names are provided:
		// Iterate over each file name.
		for _, arg := range files {
			// Open the file.
			f, err := os.Open(arg)
			if err != nil {
				// If there is an error opening the file, print the error and continue to the next file.
				fmt.Fprintf(os.Stderr, "dup02: %v\n", err)
				continue
			}
			// Count the lines from the opened file using the "countLines" function.
			countLines(f, counts)
			f.Close() // Close the file.
		}
	}

	// Iterate over the "counts" map.
	for line, n := range counts {
		// Check if the line appears more than once.
		if n > 1 {
			// Print the number of occurrences and the line.
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines counts the occurrences of each line in the given file.
func countLines(f *os.File, counts map[string]int) {
	// Create a scanner to read lines from the file.
	input := bufio.NewScanner(f)
	// Loop to read lines of text from the input.
	for input.Scan() {
		// Increment the count for the current line in the map.
		counts[input.Text()]++
	}
}
