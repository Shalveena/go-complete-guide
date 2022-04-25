package main

import "fmt"

func main() {

	age := 34

	fmt.Println(age)

	myAge := &age
	// fmt.Println(myAge)
	// fmt.Println(*myAge)

	// age = 100
	// fmt.Println(*myAge)

	// *myAge = 399
	// fmt.Println(age)
	// fmt.Println(*myAge)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(number *int) int {
	result := *number * 2
	*number = 100
	return result
}
