// Write your answer here, and then test your code.

package main

import "strings"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// censorWords() returns the original sentence with the censored words replaced with *
func censorWords(sentence string, censoredWords []string) string {

	for _, word := range censoredWords {
		sentence = strings.Replace(sentence, word, strings.Repeat("*", len(word)), -1)
	}

	return sentence
}
