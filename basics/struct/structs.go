package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// User struct to hold firstname, lastname, birthdate and createddate
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdDate time.Time
}

// outPutUserDatails method for User struct to print the user details
func (user User) outPutUserDatails() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthdate)
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var newUser *User
	
	// Get user first name
	firstName := getUserData("Please enter your first name: ")
	// Get user last name
	lastName := getUserData("Please enter your last name: ")
	// Get user birthdate
	birthdate := getUserData("Please enter your birth date (MM/DD/YYYY): ")
	
	// Create new user struct
	newUser = NewUser(firstName, lastName, birthdate)

	// Print user details
	newUser.outPutUserDatails()
}

// getUserData function to get input from user
func getUserData(promptText string) string  {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	
	// Clean the input
	claenedInput := strings.Replace(userInput, "\n", "", -1)

	return claenedInput
}

// NewUser function to create new user struct
func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		fName,
		lName,
		bDate,
		created,
	}

	return &user
}