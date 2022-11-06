package main

import "fmt"

/*
 Time to practice
 1- Create a new GO code file, set a package and add the main function;
*/
func main()  {

	// 2- Add two new variables to the file:
	// - The number PI
	pi := 3.14
	// - Circle radius
	radius := 5

	// 3- Calculate the circle circumference (2*PI*radius) and store it in a new variable.
	calculatedCircle := 2 * pi * float64(radius)
	// Output the result in the command line
	fmt.Println(calculatedCircle)

	// 4- Switch to a different way of outputting the result: Format the string to say
	// "For a radius of X, the circle circumference is Y.YY". 
	fmt.Printf("For a radius of %v, the circle circumference is %.2f \n", radius, calculatedCircle)
}