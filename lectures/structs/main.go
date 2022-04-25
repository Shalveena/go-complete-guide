package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// We will have our own custom data structure called User, which will have the data fields defined in it. This creates a blueprint of the custom type
type User struct {
	firstName   string
	lastname    string
	birthDate   string
	createdDate time.Time
}

// Receiver function (is a method of the User struct)
func (user *User) outputDetails() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastname, user.birthDate)
}

// We can also use a helper function to help us create the instance of the struct
func NewUser(fName string, lName string, bDate string) *User {
	var user User

	created := time.Now()

	user = User{
		fName,
		lName,
		bDate,
		created,
	}

	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	// Create a concrete instance of the User type, called newUser.
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Option 3: use helper function to create the instance. newUser will have all the properties and methods available in the User struct.
	newUser = NewUser(firstName, lastName, birthdate)

	// Calling outputDetails method
	newUser.outputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}
	return cleanedInput

}
