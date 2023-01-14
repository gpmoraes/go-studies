package main

import (
	"basics-go/greeting"
	"fmt"
)

func main() {
	// print the HelloText variable from the greeting package
	fmt.Println(greeting.HelloText)
	// print the Pi variable from the greeting package
	fmt.Println(greeting.Pi)
}
