package error

import "fmt"

const (
	MissingPunctuation    = "punctuation missing"
	ExtraPunctuation      = "extra punctuation"
	IncorrectPunctuation  = "incorrect punctuation"
	CapitalisationMistake = "capitalisation mistake"
	SpellingMistake       = "spelling mistake"
	IncompleteWord        = "incomplete word"
	IncorrectWord         = "incorrect word"
	MissingWord           = "missing word"
)

func PrintFullMistakes(errs []TypingError) {
	fmt.Println("Full Mistakes:")
	errorMap := map[string][]TypingError{
		SpellingMistake: {},
		IncompleteWord:  {},
		IncorrectWord:   {},
		MissingWord:     {},
	}

	for _, err := range errs {
		errList := errorMap[err.Remark]
		errList = append(errList, err)
		errorMap[err.Remark] = errList
	}

	LogErrors(errorMap)
}

func PrintHalfMistakes(errs []TypingError) {
	fmt.Println("Half Mistakes:")
	errorMap := map[string][]TypingError{
		MissingPunctuation:    {},
		ExtraPunctuation:      {},
		IncorrectPunctuation:  {},
		CapitalisationMistake: {},
	}

	for _, err := range errs {
		errList := errorMap[err.Remark]
		errList = append(errList, err)
		errorMap[err.Remark] = errList
	}

	LogErrors(errorMap)
}

func LogErrors(errMap map[string][]TypingError) {
	for k, v := range errMap {
		fmt.Println(k, ": ", len(v))
		for _, e := range v {
			fmt.Print(e.ShortString(), ", ")
		}
		fmt.Println()
		fmt.Println("-----------")
	}
}
