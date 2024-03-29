//Time to practice what you learned!

//1) Create a new Go code file, set a package and add the main function
package main

import "fmt"

func main() {

	//2) Create two new variables:
	// - For your first name
	// - and your last name

	// Create those variables in two different ways (with and without
	// explicit type assignment).

	var firstName string = "Shalveena"
	lastName := "Rohde"

	// 3) Output the two variables in the command line and execute the code file
	//   to see the output in the command line.
	fmt.Println(firstName, lastName)

	// 4) Also add two other variables:
	//   - The current year
	//   - Your birth year

	// Again, create those variables in two different ways and set the correct
	// value type (choose an appropriate type).

	var currentYear int
	currentYear = 2022

	birthYear := 1987

	// 5) Calculate the difference ("age") between the two years and
	//   store it in a new variable. Output that result in the command line.

	age := currentYear - birthYear

	fmt.Println(age)

	// 6) Overwrite the value stored in the "current year" variable with
	//   the previous value + 1 (i.e. next year). Calculate the next year,
	//   don't just change the number.
	//   Repeat step 5) with that new value (without changing any of the previous code).

	currentYear = currentYear + 1
	age = currentYear - birthYear

	fmt.Println(age)

	// Try all those steps on your own first, then have a look at my solution
	// lecture to compare your solution to mine!
}
