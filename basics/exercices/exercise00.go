package main

import "fmt"

// Time to practice:
// 1- Create a new GO code file, set a package and add the main function
func main()  {

	// 2- Declare two new variables:
	// - For your first name
	firstName  := "Guilherme"
	// - For your last name
	var lastName string = "Moraes"
	
	// Output the two variables in the command line
	fmt.Println(firstName)
	fmt.Println(lastName)

	// 4- Also add two other variables:
	// - The current year
	currenYear := 2022
	// - The birth year
	var birthYear int
	birthYear = 1991
	
	// 5- Calculate the difference ("age") between the two years and store it in a new variable.
	age := currenYear - birthYear
	
	// Output the result in the command line
	fmt.Println(age)

	// 6- Overwrite the value stored in the "current year" variable with the previous valie + 1
	currenYear = currenYear + 1
	fmt.Println(currenYear)
}
