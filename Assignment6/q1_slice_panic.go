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
