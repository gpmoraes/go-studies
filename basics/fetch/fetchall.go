package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now() // Record the current time

	ch := make(chan string) // Create a channel to receive the result strings

	urls := []string{
		"https://www.bbc.co.uk",
		"https://www.apple.com",
		"https://www.aa.com",
		"http://www.linkedin.com",
		"https://www.oracle.com",
		"https://www.docker.com",
		"https://www.github.com",
		"https://www.disneyplus.com",
		"https://www.youtube.com",
		"http://www.facebook.com",
		"https://www.medium.com",
		"https://www.cbs.com",
		"https://www.southwest.com",
		"http://www.hbo.com",
		"http://www.salesforce.com",
		"https://www.nytimes.com",
		"https://www.dominos.com",
		"https://www.wikipedia.org",
		"http://www.instagram.com",
		"http://www.stackoverflow.com",
		"https://www.spotify.com",
		"http://www.slack.com",
		"http://www.jetblue.com",
		"https://www.airbnb.com",
		"http://www.wordpress.com",
		"https://www.ebay.com",
		"http://www.paypal.com",
		"https://www.booking.com",
		"http://www.yahoo.com",
		"http://www.nbc.com",
		"http://www.dropbox.com",
		"http://www.example.com",
		"http://www.doordash.com",
		"https://www.uber.com",
		"https://www.google.com",
		"https://www.redfin.com",
		"https://www.zillow.com",
		"http://www.amazon.com",
		"https://www.adidas.com",
		"https://www.target.com",
		"https://www.lowes.com",
		"http://www.walmart.com",
		"http://www.realtor.com",
		"http://www.ikea.com",
		"https://www.opentable.com",
		"https://www.microsoft.com",
		"https://www.kayak.com",
		"http://www.lyft.com",
		"http://www.abc.com",
		"http://www.mcdonalds.com",
		"https://www.etsy.com",
		"http://www.alibaba.com",
		"http://www.nike.com",
		"http://www.homedepot.com",
		"http://www.hotels.com",
		"https://www.netflix.com",
		"http://www.trulia.com",
		"http://www.weather.com",
		"https://www.craigslist.org",
	}

	// Launch goroutines to fetch the URLs concurrently
	for _, url := range urls {
		go fetch(url, ch)
	}

	// Receive and print the result strings from the channel
	for range urls {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("%.2fs elapsed\n", elapsed.Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now() // Record the start time for this URL

	resp, err := http.Get(url) // Send an HTTP GET request to the URL
	if err != nil {
		ch <- fmt.Sprint(err) // Send error message to the channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // Read and discard the response body
	resp.Body.Close()                                 // Close the response body
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err) // Send error message to the channel
		return
	}

	elapsed := time.Since(start)                                      // Calculate the elapsed time for this URL
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed.Seconds(), nbytes, url) // Send result string to the channel
}
