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
	fmt.Printf("Original user age: %v ho", userAgeInput)

	// Clean up the user's input age
	userAgeInput = strings.Replace(userAgeInput, "\r\n", "", -1)
	fmt.Printf("Cleaned user age: %vHi \n", userAgeInput)

	// Change the age from string into integar
	age, _ := strconv.ParseInt(userAgeInput, 10, 64)
	fmt.Printf("Parsed user age: %v", age)

}
