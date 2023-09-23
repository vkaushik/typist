package evaluate

import (
	"strings"

	"github.com/vkaushik/typist/pkg/spelling"
)

type TypingError struct {
	Error    string
	Expected string
	Actual   string
}

func GetErrors(master, test string) ([]TypingError, []TypingError) {
	var fullMistakes []TypingError
	var halfMistakes []TypingError

	// add white space errors
	wsErr := GetWhiteSpaceError(master, test)
	halfMistakes = append(halfMistakes, wsErr...)

	masterWords := strings.Fields(master)
	testWords := strings.Fields(test)

	i := 0
	for len(testWords) > 0 {
		mw := masterWords[0]
		tw := testWords[0]

		res := CompareWords(mw, tw)
		if res == NoError {
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		} else if res == CapitalizationError {
			mistake := TypingError{
				Error:    GetErrorString(res),
				Expected: mw,
				Actual:   tw,
			}
			halfMistakes = append(halfMistakes, mistake)
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		} else if ContainsPunctuation(mw) || ContainsPunctuation(tw) {
			// TODO: handle punctuation scenarios here e.g. "hello. world" vs "hello.world", where candidate missed a space
			// use SplitByPunctuation
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		} else if res == IncompleteWordError {
			mistake := TypingError{
				Error:    GetErrorString(res),
				Expected: mw,
				Actual:   tw,
			}
			fullMistakes = append(fullMistakes, mistake)
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		} else if res == SpellingError {
			mistake := TypingError{
				Error:    GetErrorString(res),
				Expected: mw,
				Actual:   tw,
			}
			fullMistakes = append(fullMistakes, mistake)
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		} else if isT, fullErr, halfErr := IsTranspositionError(masterWords, testWords); isT {
			mistake := TypingError{
				Error:    "transposition error",
				Expected: masterWords[0] + " " + masterWords[1],
				Actual:   testWords[0] + " " + testWords[1],
			}
			halfMistakes = append(halfMistakes, mistake)
			halfMistakes = append(halfMistakes, halfErr...)
			fullMistakes = append(fullMistakes, fullErr...)

			masterWords = masterWords[2:]
			testWords = testWords[2:]
		} else {
			mistake := TypingError{
				Error:    GetErrorString(res),
				Expected: mw,
				Actual:   tw,
			}
			fullMistakes = append(fullMistakes, mistake)
			masterWords = masterWords[i+1:]
			testWords = testWords[i+1:]
		}
	}

	return fullMistakes, halfMistakes
}

func GetWhiteSpaceError(master, test string) []TypingError {
	var errs []TypingError
	for c := countWhitespaceErrors(master, test); c > 0; c-- {
		err := TypingError{
			Error: "whitespace ommitted or missing",
		}
		errs = append(errs, err)
	}

	return errs
}

func countWhitespaceErrors(master, test string) int {
	masterCount := spelling.CountWhitespaces(master)
	testCount := spelling.CountWhitespaces(test)

	diff := masterCount - testCount

	if diff < 0 {
		return -1 * diff
	} else {
		return diff
	}
}

// if Transposition is done with capitalisation issue, then fine
// if transposition is done with incpmolete word or spelling mistake then
func IsTranspositionError(masterWords, testWords []string) (bool, []TypingError, []TypingError) {
	var fullErrs []TypingError
	var halfErrs []TypingError

	if len(masterWords) < 2 || len(testWords) < 2 {
		return false, fullErrs, halfErrs
	}
	mw, tw := masterWords[0], testWords[0]
	nextMw, nextTw := masterWords[1], testWords[1]

	res1 := CompareWords(nextMw, tw)
	res2 := CompareWords(mw, nextTw)

	if res1 == NoError && res2 == NoError {
		return true, fullErrs, halfErrs
	}

	if !(res1 == NoError || res1 == CapitalizationError || res1 == SpellingError) && (res2 == NoError || res2 == CapitalizationError || res2 == SpellingError) {
		return false, fullErrs, halfErrs
	}

	if res1 == CapitalizationError {
		err := TypingError{
			Error:    GetErrorString(res1),
			Expected: nextMw,
			Actual:   tw,
		}
		halfErrs = append(halfErrs, err)
	}

	if res2 == CapitalizationError {
		err := TypingError{
			Error:    GetErrorString(res1),
			Expected: mw,
			Actual:   nextTw,
		}
		halfErrs = append(halfErrs, err)
	}

	if res1 == SpellingError {
		err := TypingError{
			Error:    GetErrorString(res1),
			Expected: nextMw,
			Actual:   tw,
		}
		fullErrs = append(fullErrs, err)
	}

	if res2 == SpellingError {
		err := TypingError{
			Error:    GetErrorString(res1),
			Expected: mw,
			Actual:   nextTw,
		}
		fullErrs = append(fullErrs, err)
	}

	return true, fullErrs, halfErrs
}

/*
Rules:
Following are full mistakes:
	1. For every omission of a word/figure.
	2. For every substitution of a wrong word/figure, except transposition of
	words.
	3. For every addition of a word/figure not found in the passage.
	4. For every spelling error committed by way of repetition, or addition, or
	omission, or substitution of a letter/letters, e.g. the word 'spelling'
	typed as seeplings; seplling; speling; seepling; spelling etc.
	5. For repetition of word/figure, e.g. 'I shall shall be grateful
	6. Incomplete words (half typed words will be treated as mistake).

Following errors are treated as half mistakes:
	1. Spacing Errors: Where no space is provided between two words,
	e.g. 'Ihope', or undesired space is provided between the words or
	letters of a word e.g. 'I have', 'I have' (space left between a
	word).
	2. Wrong Capitalisation: Wrong typing of a capital letter for small letter
	or vice-versa.
	(This does not apply in respect of Hindi Typewriting scripts)
	3. Punctuation Errors: Where the punctuation mark is omitted or added
	or substituted by another.
	4. Transposition Errors: Where words are transposed, e.g. the words
	'I hope' typed as 'hope I'.
	5. Paragraphic Errors: Half mistake shall be treated for each irrational
	para, where the space given before starting of any paragraph is not
	uniform, i.e paragraph given manual spaces; without pressing the
	Tab Key, will be treated half-mistake.




1. if a word is misspelled or missing, this condition have following exceptions, which will make it half mistake:
	a. If first it's capitalisation mistake
	b. if there's word transposition e.g. I hope is written as hope I
2. if an extra word is added

Following are considered as half mistakes:
1. punctuation missing
2. extra punctuation
3. wrong capitalisation
4. if paragraph is not started with a tab



*/
