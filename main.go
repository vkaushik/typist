package main

import (
	"fmt"

	"github.com/vkaushik/typist/internal/evaluate"
)

func main() {
	fmt.Println("typist test")
	fe, he := evaluate.GetErrors(masterText, testText)
	PrintErrors(fe, he)
}

var masterText = "so Hello world how are world"
var testText = "so h e.l l o world how are world"

func PrintErrors(fullErrors []evaluate.TypingError, halfErrors []evaluate.TypingError) {
	fmt.Println("Full Errors:")
	for i, err := range fullErrors {
		fmt.Printf("Error %d:\n", i+1)
		fmt.Printf("Expected: %s\n", err.Expected)
		fmt.Printf("Actual  : %s\n", err.Actual)
		fmt.Printf("Error   : %s\n\n", err.Error)
	}

	fmt.Println("Half Errors:")
	for i, err := range halfErrors {
		fmt.Printf("Error %d:\n", i+1)
		fmt.Printf("Expected: %s\n", err.Expected)
		fmt.Printf("Actual  : %s\n", err.Actual)
		fmt.Printf("Error   : %s\n\n", err.Error)
	}
}
