/*
	In the below code snippet concurrent goroutines execution corrupts a piece of data by accessing it simultaneously it leads in raise condition.
	Code snippet output when you run this : 1 is Even
	Solution snipped output once code is corrected: 0 is Even
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	var waitGroup sync.WaitGroup
	n := 0

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Printf("%d is even", n)
			return
		}
		fmt.Printf("%d is odd", n)
	}()

	// wait till the go routine is executed
	waitGroup.Wait()
	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
