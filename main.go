package main

import (
	"fmt"

	"github.com/vkaushik/typist/internal/evaluate"
)

func main() {
	fmt.Println("typist test")

	reference := "spelling"
	test := "seeplings"
	result := evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "example"
	test = "example"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "example"
	test = "exaample"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "hello"
	test = "world"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "excited"
	test = "excit"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "development"
	test = "elopment"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))

	reference = "development"
	test = "evelopment"
	result = evaluate.CompareWords(reference, test)
	fmt.Println(reference, test, evaluate.GetErrorString(result))
}
