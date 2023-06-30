/*
This Go program reads lines of text from the standard input and identifies the lines that appear more than once.
It then prints each duplicated line followed by the number of times it appears.

Algorithm:
 1. Initialize an empty map called "counts" to store the counts of each line.
 2. Create a new "Scanner" named "input" to read lines from the standard input.
 3. Enter a loop to read lines of text from the input until there are no more lines.
    3.1. Read the next line from the input using "input.Scan()".
    3.2. Increment the corresponding count in the "counts" map for the current line.
 4. After reading all the input lines, enter another loop to iterate over the "counts" map.
    4.1. For each line and its corresponding count in the "counts" map:
    4.1.1. Check if the count (number of occurrences) is greater than 1.
    4.1.1.1. If yes, print the number of occurrences and the line.
 5. End of the program.

The purpose of this program is to efficiently find and print the lines of text that appear more than once in the input, along with the number of times they appear. The map data structure is used to keep track of the counts, which allows for a fast and straightforward identification of duplicated lines.
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

	// Create a new Scanner to read from standard input.
	input := bufio.NewScanner(os.Stdin)

	// Loop to read the lines of text from input.
	for input.Scan() {
		// Increment the count for the current line in the map.
		counts[input.Text()]++
	}

	// Loop to iterate over the counts.
	for line, n := range counts {
		// Check if the line appears more than once.
		if n > 1 {
			// Print the number of occurrences and the line.
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
