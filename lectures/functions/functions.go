package main

import (
	"fmt"
	"math/rand"
)

// function to add two numbers
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

// function to print the input it receives, in a new line. Should only work with integer inputs.
func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}

// Generate random numbers
func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}

func main() {

	a, b := generateRandomNumbers()
	sum := add(a, b)
	fmt.Println(sum)

	printNumber(sum)
}
