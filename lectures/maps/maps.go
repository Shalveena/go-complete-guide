package main

import "fmt"

func main() {

	// A Map is another type of data structure.

	// Creating a new Map
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)

	// Mutating Maps

	// Extracting a value from the Map
	fmt.Println(websites["Google"]) // https://google.com

	// Adding new key-value pairs
	// A Map is always dynamic; you can always add new key-value pairs
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites) // map[Amazon Web Services:https://aws.com Google:https://google.com LinkedIn:https://linkedin.com]
	// You can also use the above way to override values.

	// Deleting key-value pairs from the Map
	delete(websites, "Google")
	fmt.Println(websites) // map[Amazon Web Services:https://aws.com LinkedIn:https://linkedin.com]

}
