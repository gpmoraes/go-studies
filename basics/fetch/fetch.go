package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := []string{"google.com"}

	// Iterate over the URLs to fetch
	for _, url := range urls {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		// Send an HTTP GET request to the URL
		resp, err := http.Get(url)
		if err != nil {
			// If an error occurs during the request, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP Requisition Status: %v\n", resp.Status)
		defer resp.Body.Close()

		file, err := os.Create("outputBuffer.txt")
		if err != nil {
			// If an error occurs while creating the file, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		// Read the response body and write it to the file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			// If an error occurs while reading the response body, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fileInfo, err := os.Stat("outputBuffer.txt")
		if err != nil {
			// If an error occurs while getting file info, print the error message and exit
			fmt.Fprintf(os.Stderr, "fetch: getting file info: %v\n", err)
			os.Exit(1)
		}

		// Print the file information
		fmt.Println("========= File Information =========")
		fmt.Printf("Name: %s\n", fileInfo.Name())
		fmt.Printf("Size: %d\n", fileInfo.Size())
		fmt.Printf("Permissions: %s\n", fileInfo.Mode())
		fmt.Printf("Last Modified: %v\n", fileInfo.ModTime())
		fmt.Printf("Is Directory: %t\n", fileInfo.IsDir())
	}
}
