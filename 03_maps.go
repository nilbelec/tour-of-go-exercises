package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for _, word := range words {
		result[word] = result[word] + 1
	}
	return result
}

// exercise03 run exercise 03
func exercise03() {
	wc.Test(wordCount)
}
