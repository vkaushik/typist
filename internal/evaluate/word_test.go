package evaluate

import (
	"fmt"
	"testing"
)

func TestCompareWords(t *testing.T) {
	type args struct {
		reference string
		test      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareWords(tt.args.reference, tt.args.test); got != tt.want {
				t.Errorf("CompareWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareWords1(t *testing.T) {
	reference := "spelling"
	test := "seeplings"
	result := fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "example"
	test = "example"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "example"
	test = "exaample"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "hello"
	test = "world"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "excited"
	test = "excit"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "development"
	test = "elopment"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))

	reference = "development"
	test = "evelopment"
	result = fmt.Println(reference, test, CompareWords(reference, test)
	fmt.Println(reference, test, fmt.Println(reference, test, GetErrorString(result))
}
