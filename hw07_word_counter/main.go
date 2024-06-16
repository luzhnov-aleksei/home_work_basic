package main

import (
	"regexp"
	"strings"
)

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)

	re := regexp.MustCompile(`[\p{L}\p{N}]+`)

	words := re.FindAllString(text, -1)

	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount
}
