package tokenize

import (
	"regexp"
)

// SplitStringByWhitespaceOrPunctuation splits a string by whitespace or punctuation.
func SplitStringByWhitespaceOrPunctuation(input string) []string {
	// Define a regular expression pattern to match whitespace and punctuation.
	regexPattern := `[[:space:]|[:punct:]]+`

	// Compile the regular expression.
	regex := regexp.MustCompile(regexPattern)

	// Use the regular expression to split the input string.
	words := regex.Split(input, -1)

	// Remove empty strings from the result.
	var cleanedWords []string
	for _, word := range words {
		if word != "" {
			cleanedWords = append(cleanedWords, word)
		}
	}

	return cleanedWords
}
