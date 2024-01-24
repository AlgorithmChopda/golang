/*
	1. Given a string containing conversation between alice and bob. In the string, if it reaches $, it means end of alice message and if it reaches #, it means end of bob's message. and if it reaches ^,
	it means end of conversation ignore string after that.

	Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example.

	write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

	Note: there is a single space before and after colon(:) and no space before and after comma.

	e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
	output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?
*/

package main

import (
	"fmt"
)

func readConversation(conversation string, aliceChan chan<- string, bobChan chan<- string, signalChan chan<- struct{}) {
	var current string
	for _, ch := range conversation {
		switch ch {
		case '$':
			aliceChan <- current
			current = ""
		case '#':
			bobChan <- current
			current = ""
		case '^':
			signalChan <- struct{}{}
		default:
			current += string(ch)
		}
	}
}

func writeConversation(aliceChan <-chan string, bobChan <-chan string, signalChan <-chan struct{}) {
	for {
		select {
		case message := <-aliceChan:
			fmt.Printf("\nAlice : %s", message)
		case message := <-bobChan:
			fmt.Printf("\nBob : %s", message)
		case _ = <-signalChan:
			// if signal chan then return
			return
		}
	}
}

func main() {
	var conversation string
	fmt.Printf("Enter Input : ")
	fmt.Scanln(&conversation)

	aliceChan := make(chan string)
	bobChan := make(chan string)
	signalChan := make(chan struct{})
	defer func() {
		close(aliceChan)
		close(bobChan)
	}()

	go readConversation(conversation, aliceChan, bobChan, signalChan)
	writeConversation(aliceChan, bobChan, signalChan)
}
