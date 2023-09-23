package main

import (
	"fmt"
	"strings"
	"github.com/vkaushik/typist/pkg/spelling"
)

// TypingMistakes represents a collection of typing mistakes.
type TypingMistakes struct {
	FullMistakes []string
	HalfMistakes []string
}

// FindTypingMistakes finds typing mistakes in the test string compared to the master string.
func FindTypingMistakes(master, test string) TypingMistakes {
	fullMistakes := make([]string, 0)
	halfMistakes := make([]string, 0)

	// Split the master and test strings into words.
	masterWords := strings.Fields(master)
	testWords := strings.Fields(test)

	// Initialize variables to keep track of transposition errors.
	transposing := false
	transpositionWord := ""

	// Iterate through the words.
	for i := 0; i < len(masterWords) || i < len(testWords); i++ {
		if i >= len(masterWords) {
			fullMistakes = append(fullMistakes, testWords[i])
		} else if i >= len(testWords) {
			fullMistakes = append(fullMistakes, masterWords[i])
		} else {
			// Check for transposition errors.
			if !transposing && i+1 < len(testWords) && masterWords[i] == testWords[i+1] {
				transposing = true
				transpositionWord = masterWords[i]
			} else if transposing && masterWords[i] == transpositionWord {
				// Check if it's a transposition error.
				transposing = false
				fullMistakes = append(fullMistakes, transpositionWord, testWords[i])
				i++ // Skip the next word in the test string.
			} else {
				// Check if it's a spelling mistake or other full mistake.
				if spelling.IsSpellingMistake(masterWords[i], testWords[i]) {
					fullMistakes = append(fullMistakes, testWords[i])
				}
			}
		}
	}

	// Check for half mistakes.
	// You need to implement the logic for half mistakes based on the provided rules.

	return TypingMistakes{
		FullMistakes: fullMistakes,
		HalfMistakes: halfMistakes,
	}
}

func main() {
	master := "this is master text for everyone text master"
	test := "this text master for everyone text master"

	// Find typing mistakes.
	mistakes := FindTypingMistakes(master, test)

	// Print full mistakes.
	fmt.Println("Full Mistakes:")
	for _, mistake := range mistakes.FullMistakes {
		fmt.Println("-", mistake)
	}

	// Print half mistakes.
	fmt.Println("\nHalf Mistakes:")
	for _, mistake := range mistakes.HalfMistakes {
		fmt.Println("-", mistake)
	}
}
