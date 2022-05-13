package main

import "fmt"

// import "fmt"

func main() {
	/*
		// Creating an array with 4 elements/values in it. When you use this syntax, you need to supply the element data as well (the data stored in the array)
		prices := [4]float64{10.99, 9.99, 45.05, 20.00}

		// Creating the array without having to input the data to be stored in the array. You should add in a number in the square brackets to indicate the length of the array, otherwise it will be an empty array with no items at all!
		var productNames [4]string
		fmt.Println(prices)       // [10.99 9.99 45.05 20]
		fmt.Println(productNames) // [    ]

		// Accessing a specific element in an array
		fmt.Println(prices[1]) // note: dot notation not available for arrays! Outputs 9.99

		// Setting value for specific elements in the array
		productNames[2] = "A Carpet"
		productNames[0] = "A Pool"

		fmt.Println(productNames)

		// Slices
		featuredProducts := prices[1:3] // including the element at position 1 and excluding the element at position 3
		fmt.Println(featuredProducts)   // [9.99 45.05]

		featuredProducts = prices[:3] // You get a slice of the array from the first element up to (but exclucing) the element at position 3
		fmt.Println(featuredProducts) // [10.99 9.99 45.05]

		// Can also do the opposite: start at a specific element and go all the way to the end of the original array
		featuredProducts = prices[1:]
		fmt.Println(featuredProducts) // [9.99 45.05 20]

		// Note: can't use negative indices in the square bracket. E.g. [-1]

		// Slices can also be created based on other slices
		highlightedPrices := featuredProducts[:1] // [9.99]
		fmt.Println(highlightedPrices)
	*/

	/*
		// changes to the slice are also changes to the original array
		prices := [5]float64{10.99, 9.99, 45.05, 20.00, 52.00}
		fmt.Println("prices:", prices)

		featuredPrices := prices[1:4] // [9.99 45.05 20]
		//highlightedPrices := featuredPrices[1:] // [45.05 20]
		//highlightedPrices[0] = 100

		//fmt.Println(highlightedPrices) // [100 20]
		fmt.Println("featured prices:", featuredPrices) // [9.99 45.05]
		//fmt.Println(prices) // [10.99 9.99 100 20]

		fmt.Println("length:", len(featuredPrices), "capacity:", cap(featuredPrices)) // 3 3

		highlightedPrices := featuredPrices[:2] // [9.99]
		fmt.Println("highlighted prices", highlightedPrices)
		fmt.Println("length:", len(highlightedPrices), "capacity:", cap(highlightedPrices))

		highlightedPrices = highlightedPrices[:4]
		fmt.Println("highlighted prices", highlightedPrices)
		fmt.Println("length:", len(highlightedPrices), "capacity:", cap(highlightedPrices))
		//fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1 3

		// featuredPrices = featuredPrices[:4]
		// fmt.Println(featuredPrices)
	*/

	// Creating Dynamic Lists With Slices
	numbers := []int{9, 47, 6, 51, 77}
	numbersSliced := numbers[0:3]
	fmt.Println(numbersSliced)

	updatedNumbers := append(numbersSliced, 8)
	fmt.Println("numbersSliced: ", numbersSliced)
	fmt.Println("updatedNumbers: ", updatedNumbers)
	fmt.Println("numbers: ", numbers)

}
