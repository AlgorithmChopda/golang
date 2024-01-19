/*

	1. The given program accepts an integer value between 1 to 4 from the user.
	There is an option associated with each value, which is basically objects of different data types with some associated value.

	Write a method named "AcceptAnything" which takes any type of data as input.

	Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

	integer :
	"This is a value of type Integer, <value>"

	string :
	"This is a value of type String, <value>"

	boolean :
	"This is a value of type Boolean, <value>"

	Custom data type Hello :
	"This is a value of type Hello, <value>"
*/

package main

import "fmt"

type Student struct {
	name string
	age  int
}

// using any for accepting all type of value(input)
func AcceptAnything(value any) {
	switch value := value.(type) {
	case int:
		fmt.Printf("This is a value of type Integer: %#v", value)
	case string:
		fmt.Printf("This is a value of type String: %#v", value)
	case bool:
		fmt.Printf("This is a value of type Boolean: %#v", value)
	case Student:
		fmt.Printf("This is a value of type Student: %#v", value)
	default:
		fmt.Printf("No such type found")
	}
}

func processChoice(choice int) {
	switch choice {
	case 1:
		// int
		AcceptAnything(100)
	case 2:
		// string
		AcceptAnything("Hello Golang")
	case 3:
		// boolean
		AcceptAnything(true)
	case 4:
		// type Student
		var student Student
		AcceptAnything(student)
	}
}

func main() {
	var choice int
	fmt.Printf("Enter a no between 1 to 4: ")
	fmt.Scanln(&choice)

	processChoice(choice)
}
