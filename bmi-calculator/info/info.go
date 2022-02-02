package info

import "fmt"

const welcomeText string = "Welcome to the BMI Calculator"
const divider = "_______________________________"
const WeightPrompt = "Please enter your weight (kg): "
const HeightPrompt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(welcomeText)
	fmt.Println(divider)
}
