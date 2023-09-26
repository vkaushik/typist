package tokenize

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func IsAlphanumeric(input string) bool {
	if len(input) != 1 {
		return true
	}
	in := []rune(input)[0]

	return unicode.IsLetter(in) || unicode.IsDigit(in)
}

func PrintWordWithIndex(words []string) {
	fmt.Println()
	for i, w := range words {
		fmt.Printf(`{%d: "%s"}, `, i, w)
	}
	fmt.Println()
}

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

func Tokenize(text string) []string {
	isWord := false
	var word strings.Builder
	var words []string

	for _, ch := range text {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) { // a letter of number
			isWord = true
			word.WriteRune(ch)
		} else { // punctuation or whitespace
			if isWord { // inside word
				words = append(words, word.String())
				word.Reset()
			}

			words = append(words, string(ch))
			isWord = false
		}
	}

	return words
}

func splitWordsAndSpaces1(input string) []string {
	// Define a regular expression to match words and spaces
	re := regexp.MustCompile(`[\p{L}\p{N}]+|[\s]+`)

	// Use FindAllString to find all matches in the input
	result := re.FindAllString(input, -1)

	return result
}

func splitWordsAndSpaces(input string) []string {
	var result []string

	for _, char := range input {
		if unicode.IsSpace(char) {
			result = append(result, " ")
		} else {
			result = append(result, string(char))
		}
	}

	return result
}
