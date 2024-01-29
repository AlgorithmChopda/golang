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
