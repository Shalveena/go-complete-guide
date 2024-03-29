// Time to practice what you learned!

// 1) Create a new Go code file and initialize it correctly.

package main

import "fmt"

func main() {

	// 2) Add two new variables to the file:

	// - The number PI (3.14)
	// - A circle radius with a value of 5
	PI := 3.14
	radius := 5

	// 3) Calculate the circle circumference (2*PI*radius) and store it in a new variable.
	circumference := 2 * PI * float64(radius)
	//   Output that value in the command line. Make sure that the value is correct!
	fmt.Println(circumference)
	//   Side-note: The result may have a tiny precision issue - you can ignore that.

	// 4) Switch to a different way of outputting the result: Format the string to say
	//   "For a radius of X, the circle circumference is Y.YY". X and Y should be concrete values.
	//   Hint: You might need to look into https://golang.org/pkg/fmt/#hdr-Printing to learn how to
	//   output two decimal places (or scroll to the bottom of this file).
	fmt.Printf("For a radius of %v, the circle's circumference is %.2f.", radius, circumference)

	// You can format a floating point number to be output as Y.YY (two decimal places) via
	// the %.2f placeholder in a formatted string.
}
