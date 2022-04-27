package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// 1. Create a new type / data structure
type Product struct {
	id          string
	title       string
	description string
	price       float64
}

// 2(b) Creation helper function
func NewProduct(ID string, title string, dscript string, price float64) *Product {

	return &Product{
		ID,
		title,
		dscript,
		price,
	}

}

// 3(a). Add a "connected function" that outputs the data
func (product *Product) outputData() {

	fmt.Printf("The product with the ID %v is %v, which is %v. The price of this item is %v \n", product.id, product.title, product.description, product.price)
}

// 5) Bonus: Add a connected "store" function that writes that data into a file. The filename should be the unique id, the function should be called at the end of main.
func (product *Product) writeData() {

	file, _ := os.Create(product.id + ".txt")
	contentToWrite := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: %v\n",
		product.id,
		product.title,
		product.description,
		product.price,
	)
	file.WriteString(contentToWrite)

	file.Close()

}

// 4. Change the program to fetch user input values for the different data fields
func getProduct() *Product {

	fmt.Println("Please enter the following data to register a new product:")
	fmt.Println("----------------------------------------------------------")

	productId := getUserInput("Enter product ID: ")
	productTitle := getUserInput("Enter product title: ")
	productDescription := getUserInput("Enter product description: ")
	productPriceString := getUserInput("Enter product price: ")
	productPriceFloat, _ := strconv.ParseFloat(productPriceString, 64)

	return NewProduct(productId, productTitle, productDescription, productPriceFloat)
}

func getUserInput(userPrompt string) (userInput string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(userPrompt)
	userInput, _ = reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}
	return
}

func main() {

	// 4. ...and create only one concrete instance with that entered data
	firstProduct := getProduct()
	firstProduct.outputData()
	firstProduct.writeData()

	/*
		// 2(a). Create a concrete instance of this data type directly inside of the main function
		firstProduct := Product{
			"001",
			"Atomic Habits",
			"Book about creating habits",
			27.99,
		}

		// 2(b). Create a concrete instance of this data type inside of main using a "creation helper function"
		secondProduct := NewProduct("002",
			"Five Love Languages",
			"Book about relationships",
			21.99)

		firstProduct.outputData()
		secondProduct.outputData()
	*/
}

// Your Tasks
//
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
