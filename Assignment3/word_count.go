/*
	3. Word Count

	Write a Program to fulfil the following conditions.
	Input: A string containing words separated by space
	Output: A slice containing words with the highest frequency in the input string.

	Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
	Condition: Words are case-sensitive. Joe is not the same as joe.
	Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.

	Example Test Case 1:

	Input: My name is Joe and My Father's name is also Joe
	Output: [My name is Joe]
	Here, the words "My", "name", "is", "Joe" appeared 2 times each, which is also the highest frequency of any word.
	Hence the output contains all 4 words.
	Example Test Case 2:

	Input: Europe is far but the US is farther
	Output: [is]
	Here, the word "is" appeared twice which is also the highest frequency of any word.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getHighestFrequencyWords(str string) []string {
	// Trim to remove \n from end
	str = strings.Trim(str, "\n")

	var wordFrequency = map[string]int{}
	tokens := strings.Split(str, " ")

	// count frequency of each word
	maxFrequency := 0
	for _, word := range tokens {
		wordFrequency[word]++
		maxFrequency = max(maxFrequency, wordFrequency[word])
	}

	var output []string
	for _, word := range tokens {
		if wordFrequency[word] == maxFrequency {
			output = append(output, word)
			// once word is print ignore further occurences
			wordFrequency[word] = -1
		}
	}

	return output
}

func main() {
	// bufio used to scan new entire line
	fmt.Printf("Enter a Sentence : ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	sentence = strings.Trim(sentence, "\n")

	fmt.Printf("Output: %s", getHighestFrequencyWords(sentence))
}
