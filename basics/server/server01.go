package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register the handler function for the root URL ("/").
	// This function will be called whenever a request is made to the root URL.
	http.HandleFunc("/", handler)

	// Start the HTTP server and listen on port 8000.
	// The server will use the default server mux and handle all incoming requests.
	// If an error occurs, log.Fatal will log the error message and exit the program.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler is the function that handles requests to the root URL ("/").
// It takes an http.ResponseWriter and an http.Request as parameters.
// The http.ResponseWriter is used to write the response back to the client,
// and the http.Request contains information about the incoming request.
func handler(w http.ResponseWriter, r *http.Request) {
	// Use the fmt.Fprintf function to write the response to the http.ResponseWriter.
	// In this case, it writes the formatted string "URL.Path = %q\n" to the response,
	// with r.URL.Path as the value to be formatted.
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
