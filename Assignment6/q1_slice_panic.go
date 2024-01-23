/*
	1. In the given code, the accessSlice function accepts a slice and index.
	If the value is present in the slice at that index, the program should print the following.

	"item <index>, value <value at that index in slice>"

	But if the index does not hold any value,
	it will lead to an index out of range panic in our program.

	Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
	Also, Print the following after handling panic

	"internal error: <recovered panic value>"

	Example Test Case 1 :

	Input: 3
	Output:
	item 3, value 6
	Testing Panic and Recover
*/

package main

import "fmt"

func accessSlice(index int, slice []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Internal error: %v", r)
		}
	}()

	fmt.Printf("Item %d, Value %d", index, slice[index])
}

func main() {
	var slice = []int{10, 20, 30, 40, 50}
	var index int
	fmt.Printf("Enter index : ")
	fmt.Scanln(&index)

	accessSlice(index, slice)
}
