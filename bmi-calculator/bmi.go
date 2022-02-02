package main

import (
	"github.com/shalveena/bmicalculator/info"
)

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func main() {

	// Welcome message/intro
	info.PrintWelcome()

	// Get user input of weight and height
	height, weight := getUserWeightAndHeight()

	// Calculate BMI
	bmi := calculateBMI(weight, height)

	// Output BMI
	printBMI(bmi)
}
