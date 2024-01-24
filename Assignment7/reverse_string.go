/*
	2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched

	e.g: Input: test123 output: 321tset 2
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func reverseString(input string, wg *sync.WaitGroup) {
	defer wg.Done()

	var reversedString string
	for i := len(input) - 1; i >= 0; i-- {
		reversedString += string(input[i])
	}
	fmt.Printf("Reversed string : %s", reversedString)
	fmt.Printf("\nGo routines running : %d", runtime.NumGoroutine())
}

func main() {
	var input string
	fmt.Printf("Enter String : ")
	fmt.Scanln(&input)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go reverseString(input, &wg)
	wg.Wait()
}
