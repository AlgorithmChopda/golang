/*
	1. Find the Index of the element

	Write a program to store the day(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) against the respective index of the day(1, 2, 3, 4, 5, 6, 7) in a map.
	You can hard-code the map in your program.
	Take an integer input from the user and print the day stored against that index and if nothing is to be found for that index,
	Print "Not a day"

	Hint: Following code can be used to take an integer input from the user in the GO Programming Language
	var index int
	fmt.Scanln(&index)

	Example Test Case:
	Input: 5
	Output: Friday

	Input: 11
	Output: Not a day
*/

package main

import "fmt"

func getDayFromIndex(index int) string {
	var weekday_map = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thrusday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	day, isPresent := weekday_map[index]
	if isPresent {
		return day
	}

	return "Not a day"
}

func main() {
	var index int
	fmt.Printf("Enter the Integer : ")
	fmt.Scanln(&index)

	fmt.Printf("Output: %s", getDayFromIndex(index))
}
