package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shalveena/bmicalculator/info"
)

func main() {

	// Welcome message/intro
	fmt.Println(info.WelcomeText)
	// Get user input of weight and height
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Clean up the user input and parse the string to float64
	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)

	heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculate BMI
	bmi := weight / (height * height)

	// Output BMI
	fmt.Printf("Your BMI is %.2f", bmi)

}
