package main

import (
	"fmt"

	"github.com/Shalveena/go-complete-guide/greeting"
)

func main() {

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)

	// Working with "int" types:
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = luckyNumber / 3

	fmt.Println(evenMoreLuckyNumber)

	var newNumber = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var firstByte byte = '9'
	var secondByte byte = 9

	fmt.Println(firstByte)
	fmt.Println(secondByte)

	firstName := "Shalveena"
	lastName := "Rohde"
	// fullname := firstName + " " + lastName
	fullname := fmt.Sprint(firstName, " ", lastName)
	age := 34

	// Formatting strings
	fmt.Printf("Hi, my name is %v and I am %v years old.", fullname, age)
	// fmt.Printf("Hi, my name is %v and I am %v (Type: %T) years old.", fullname, age, age)

	fmt.Println(pi)
}

// We often use third party code packages that provide certain functionality to us. Go has a standard library, which is a "built in module" full of core packages and functionalities, which we don't need to download and install separately - e.g. fmt => when we import fmt, we are importing one of these built in packages.
// We are then using the fmt package when we do fmt.Println => use are using the Println function, which is made available by the fmt package.
