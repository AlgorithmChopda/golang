/*
	4. Return the slices
	Complete the program to return 3 slices of a given array, based on the following conditions.
	Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]

	Hint: Hard-code the given array into your program.

	Input: Two space-separated integers representing index1 and index2.
	Output: Output will contain 3 slices
	1. slice containing all the elements from the start of the array to Index1
	2. slice containing all the elements from index1 to index2
	3. slice containing all the elements from index2 to the end of the array

	Conditions to Handle:
	If Either of the input indexes is out of range of the given array, print "Incorrect Indexes" and exit the program
	Example Test Case 1:

	Input: 2 4
	Output:
	[qwe wer ert]
	[ert rty tyu]
	[tyu yui uio iop]
	Example Test Case 2:

	Input: 2 8
	Output: Incorrect Indexes
*/

package main

import (
	"fmt"
	"log"
)

func getSlices(startIndex, endIndex int) ([]string, []string, []string) {
	ar := [...]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	// check for invalid data input
	if startIndex < 0 || endIndex >= len(ar) || startIndex > endIndex {
		// Exit as per question
		log.Fatal("Incorrect Indexes")
	}

	startSlice := ar[:startIndex+1]
	middleSlice := ar[startIndex : endIndex+1]
	endSlice := ar[endIndex:]

	return startSlice, middleSlice, endSlice
}

func main() {
	var startIndex, endIndex int
	fmt.Printf("Enter start and end Index :")
	fmt.Scanln(&startIndex, &endIndex)

	startSlice, middleSlice, endSlice := getSlices(startIndex, endIndex)

	fmt.Printf("\nStarting : %s", startSlice)
	fmt.Printf("\nMiddle : %s", middleSlice)
	fmt.Printf("\nEnding : %s", endSlice)
}
