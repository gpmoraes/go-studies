package main

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/gpmoraes/bmi/info"
)

func main() {
	// header
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	
	// ask user for the information
	fmt.Print(info.WeightText)
	// as the reader will return two values we can ignore the second one
	// adding a underline at the second variable
	inputWeight, _ := reader.ReadString('\n')

	fmt.Print(info.HeightText)
	inputHeight, _ := reader.ReadString('\n')

	// clear the input line break
	inputWeight = strings.Replace(inputWeight, "\n", "", -1)
	inputHeight = strings.Replace(inputHeight, "\n", "", -1)

	// turn the input (string) into float
	weight, _ := strconv.ParseFloat(inputWeight, 64)
	height, _ := strconv.ParseFloat(inputHeight, 64)

	// calculate BMI
	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f \n", bmi)
}