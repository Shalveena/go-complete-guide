package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Create an app to calculate different things based on user input
var reader = bufio.NewReader(os.Stdin)

func main() {

	userSelection, err := getUserInput()
	// If there is an error, exit the program.
	if err != nil {
		fmt.Println("Invalid choice. Program exited.")
		return
	}

	// If there is no error, depending on the user's choice call relevant function
	switch userSelection {
	case "1":
		addUpToNumberX(int(chosenNumber))
	case "2":
		buildFactorial()
	case "3":
		sumManually()
	default:
		sumListOfNumbers()
	}
}

func getUserInput() (string, error) {

	// Prompt user input
	fmt.Println("Please choose an option (1, 2, 3, or 4)")
	fmt.Println("1) Add up all the numbers up to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of numbers")

	// Read user input
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// Clean user input
	userInput = cleanUserInput(userInput)

	// Check that user has only entered 1, 2, 3, or 4
	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("invalid input")
	}

}

func cleanUserInput(userInput string) string {
	os := runtime.GOOS
	if os == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return userInput
}

func getInputNumber() (int, error) {
	fmt.Println("Please enter a number")
	chosenNumberString, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	chosenNumberString = cleanUserInput(chosenNumberString)
	chosenNumberInt, err := strconv.ParseInt(chosenNumberString, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenNumberInt), nil
}

func addUpToNumberX() {

	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	sum := 0

	for i := 1; i <= chosenNumber; i++ {
		sum += i
		fmt.Println(sum)
	}

	return sum

}

func buildFactorial() {

}

func sumManually() {

}

func sumListOfNumbers() {

}
