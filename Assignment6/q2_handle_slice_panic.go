/*
	2. In the given code, the accessSlice function accepts a slice and index.
	If the value is present in the slice at that index, the program should print the following.

	"item <index>, value <value present at the index>"

	But if the index does not hold any value,
	it will lead to an index out of range panic in our program.
	So in order to safeguard our program from panicking, add a condition to handle the condition if

	index > lengthOfSlice - 1

	and return an error from the accessSlice function if the above condition is met.
	The error message should be the following :

	"length of the slice should be more than index"

	Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message.

	Example Test Case 1 :

	Input: 3
	Output:
	item 3, value 6
*/

package main

import (
	"errors"
	"fmt"
)

func accessSlice(index int, slice []int) error {
	length := len(slice)
	if index > length-1 {
		return errors.New("length of the slice should be more than index")
	}

	fmt.Printf("Item %d, Value %d", index, slice[index])
	return nil
}

func main() {
	var slice = []int{10, 20, 30, 40, 50}
	var index int
	fmt.Printf("Enter index : ")
	fmt.Scanln(&index)

	err := accessSlice(index, slice)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
}
