# Go User Struct

This is a Go program that creates a new user struct and prompts the user to enter their first name, last name, and birthdate. The program then calls the `outPutUserDatails` method on the new user struct which prints out the user's name and birthdate.

## Functions

### `main()`

The `main` function uses the `getUserData` function to prompt the user for their information, and then calls the `NewUser` function to create a new user struct with the provided information. The new user struct is then stored in the `newUser` variable, which is passed to the `outPutUserDatails` method to print the user's information.

### `getUserData(promptText string) string`

The `getUserData` function takes in a `promptText` argument, and uses it to prompt the user for input. It then uses the bufio package's Reader to read the user's input and replaces the newline character at the end of the input with an empty string using the strings package's Replace function.

### `NewUser(fName string, lName string, bDate string) *User`

The `NewUser` function takes in three arguments: `fName`, `lName`, and `bDate`, which represent the user's first name, last name, and birthdate respectively. It creates a new `time.Time` object called `created`, which is the current time, and then creates a new User struct with the provided information and the created time. It then returns the address of the newly created user struct.

### `outPutUserDatails()`

The `outPutUserDatails` method is a method of the `User` struct. It takes no arguments and prints out the user's first name, last name, and birthdate using the fmt package's `Printf` function.

## Structs

### `User`

The `User` struct holds the following fields:
- `firstName`: a string representing the user's first name
- `lastName`: a string representing the user's last name
- `birthdate`: a string representing the user's birthdate
- `createdDate`: a `time.Time` object representing the date the user struct was created

## Packages

This program uses the following packages:
- `bufio`: for reading user input
- `fmt`: for printing output
- `os`: for accessing the operating system's input/output
- `strings`: for manipulating strings
- `time`: for getting the current time
