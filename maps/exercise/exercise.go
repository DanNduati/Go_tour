package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// hint: You might find strings.Fields helpful
	// Create a map to store the count
	word_count := make(map[string]int)
	// get a slice of all the words in the string
	words := strings.Fields(s)
	// loop through all the words in the slice
	for _, word := range words {
		word_count[word]++
	}
	return word_count
}

func main() {
	// Exercise: Implement WordCount
	wc.Test(WordCount)
}
