package spelling

import "strings"

// CountWhitespaces counts the total number of whitespaces in a string.
func CountWhitespaces(input string) int {
	count := 0

	// Iterate through the characters in the string.
	for _, char := range input {
		// Check if the character is a whitespace.
		if strings.ContainsRune(" \t\n\r\v\f", char) {
			count++
		}
	}

	return count
}
