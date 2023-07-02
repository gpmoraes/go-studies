package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//urls := []string{"https://example.com", "https://google.com", "https://github.com"}
	urls := []string{"https://google.com"}

	// Iterate over the command-line arguments starting from index 1
	// Each argument represents a URL to fetch
	for _, url := range urls {
		// Send an HTTP GET request to the URL
		resp, err := http.Get(url)
		if err != nil {
			// If an error occurs during the request, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Read the response body into a byte slice
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			// If an error occurs while reading the response body, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// Print the response body to the standard output
		fmt.Printf("%s\n", b)
	}
}
