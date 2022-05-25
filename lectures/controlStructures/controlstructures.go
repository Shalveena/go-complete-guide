package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Prompt user to enter age
	fmt.Print("Please enter your age: ")

	// Save user's age to a variable
	reader := bufio.NewReader(os.Stdin)
	userAgeInput, _ := reader.ReadString('\n')
	// fmt.Printf("Original user age: %v ho", userAgeInput)

	// Clean up the user's input age
	userAgeInput = strings.Replace(userAgeInput, "\r\n", "", -1)
	// fmt.Printf("Cleaned user age: %vHi \n", userAgeInput)

	// Change the age from string into integar
	userAge, _ := strconv.ParseInt(userAgeInput, 10, 64)
	// fmt.Printf("Parsed user age: %v", userAge)

	// If Statement
	if userAge >= 30 && userAge < 60 {
		fmt.Println("Congratulations, you're eligible to apply")
	} else if userAge >= 60 || userAge < 18 {
		fmt.Println("Welcome to the Golden Age!")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the working club")
	} else {
		fmt.Println("Sorry, you're not old enough 😐")
	}

}
