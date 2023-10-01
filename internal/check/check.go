package check

import (
	"fmt"
	"strings"

	"github.com/vkaushik/typist/pkg/error"
	"github.com/vkaushik/typist/pkg/spelling"
	"github.com/vkaushik/typist/pkg/tokenize"
)

func GetErrors(master, test string) ([]error.TypingError, []error.TypingError) {
	var fullMistakes []error.TypingError
	var halfMistakes []error.TypingError

	// add white space errors

	if len(master) <= len(test) {
		master = master + master
	}

	// split strings by spaces and punctuations. space and punctuation also there in resulting list.
	masterWordList := tokenize.Tokenize(master)
	testWordList := tokenize.Tokenize(test)
	WordIndexInMaster := -1
	tokenize.PrintWordWithIndex(masterWordList) // for easy debugging

	for len(testWordList) > 0 && len(masterWordList) > 0 { // compare words from master text one by one
		WordIndexInMaster++
		masterWord := masterWordList[0]
		testWord := testWordList[0]

		fmt.Printf(`"%s" == "%s"`, masterWord, testWord) // for easy debugging
		fmt.Println()

		if masterWord == testWord { // no mistake
			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			continue
		}

		if !tokenize.IsAlphanumeric(masterWord) && !tokenize.IsAlphanumeric(testWord) {
			err := error.NewTypingError(error.IncorrectPunctuation, masterWord, testWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)
			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			continue
		}

		if !tokenize.IsAlphanumeric(masterWord) { // if it is space or punctuation, that means user missed it
			err := error.NewTypingError(error.MissingPunctuation, masterWord, testWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)

			masterWordList = masterWordList[1:]
			fmt.Println(err.Error())
			continue
		}

		if !tokenize.IsAlphanumeric(testWord) { // if masterWord is alphanumeric but testWord is space of punctuation, that means that's an extra punctuation from user
			err := error.NewTypingError(error.ExtraPunctuation, masterWord, testWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)

			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		}

		// here onwards both test and master word are alphanumeric word

		if strings.EqualFold(masterWord, testWord) { // Check for capitalisation mistake
			err := error.NewTypingError(error.CapitalisationMistake, masterWord, testWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)

			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		}

		// did user typed a punctuation by mistake i.e. get next testWord and join it with testWord compare against masterWord
		if len(testWordList) < 2 { // no more words in master list
			err := error.NewTypingError(error.IncorrectWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			masterWordList = masterWordList[1:]
			continue
		}

		// find the next test word to figure out if user missed the word completely
		var nextTestWord string
		var punctuationMissingMistakesIfMissingWordCaseInTest []error.TypingError // TODO: Where is it used
		var j int
		for j = 1; j < len(testWordList); j++ {
			nextTestWord = testWordList[j]
			if tokenize.IsAlphanumeric(nextTestWord) {
				break
			}
			err := error.NewTypingError(error.MissingPunctuation, masterWord, nextTestWord, WordIndexInMaster)
			punctuationMissingMistakesIfMissingWordCaseInTest = append(punctuationMissingMistakesIfMissingWordCaseInTest, err)
		}

		if j == len(testWordList) { // no alphanueric word found in remaining test word list
			err := error.NewTypingError(error.IncorrectWord, "", testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			masterWordList = []string{}
			continue
		}

		if strings.EqualFold(masterWord, testWord+nextTestWord) {
			err := error.NewTypingError(error.ExtraPunctuation, masterWord, nextTestWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)
			testWordList = testWordList[j+1:]
			masterWordList = masterWordList[1:]
			continue
		}

		// now either user have missed the word
		// or typed an incorrect word - possible mistakes

		if len(masterWordList) < 2 { // no more words in master list
			err := error.NewTypingError(error.IncorrectWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			masterWordList = masterWordList[1:]
			continue
		}

		// find the next master word to figure out if user missed the word completely
		var nextMasterWord string
		var punctuationMissingMistakesIfMissingWordCase []error.TypingError
		var i int
		for i = 1; i < len(masterWordList); i++ {
			nextMasterWord = masterWordList[i]
			if tokenize.IsAlphanumeric(nextMasterWord) {
				break
			}
			err := error.NewTypingError(error.MissingPunctuation, nextMasterWord, testWord, WordIndexInMaster)
			punctuationMissingMistakesIfMissingWordCase = append(punctuationMissingMistakesIfMissingWordCase, err)
		}

		if i == len(masterWordList) { // no alphanueric word found in remaining master word list
			err := error.NewTypingError(error.IncorrectWord, "", testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			masterWordList = []string{}
			continue
		}

		var userMissedWord bool
		if nextMasterWord == testWord { // testWord is matching with nextMasterWord i.e. user missed the masterWord
			userMissedWord = true
		} else if spelling.IsSpellingMistake(nextMasterWord, testWord) {
			userMissedWord = true
			err := error.NewTypingError(error.SpellingMistake, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
		} else if strings.EqualFold(nextMasterWord, testWord) {
			userMissedWord = true
			err := error.NewTypingError(error.CapitalisationMistake, masterWord, testWord, WordIndexInMaster)
			halfMistakes = append(halfMistakes, err)
		}

		if userMissedWord {
			fmt.Println("user missed the word", masterWord, testWord, WordIndexInMaster)
			err := error.NewTypingError(error.MissingWord, masterWord, "", WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			halfMistakes = append(halfMistakes, punctuationMissingMistakesIfMissingWordCase...)

			if i > len(masterWordList)-2 {
				masterWordList = []string{}
				continue
			}
			masterWordList = masterWordList[i+1:]
			testWordList = testWordList[1:]
			continue
		} else if strings.HasPrefix(masterWord, testWord) || strings.HasPrefix(testWord, masterWord) { // Check for incomplete words (half-typed).
			// check if prefix issue is because of a punctuation missing
			compositeWord := masterWord + nextMasterWord
			if strings.EqualFold(compositeWord, testWord) {
				err := error.NewTypingError(error.MissingPunctuation, masterWord, testWord, WordIndexInMaster)
				halfMistakes = append(halfMistakes, err)
				masterWordList = masterWordList[i+1:]
				testWordList = testWordList[1:]
				continue
			}

			compositeTestWord := testWord + nextTestWord
			if strings.EqualFold(compositeWord, compositeTestWord) { // and the --> an dthe  // TODO: in case multiple extra non-alphanumeric chars in test or capitalisation issue
				err := error.NewTypingError(error.MissingPunctuation, masterWord, testWord, WordIndexInMaster)
				halfMistakes = append(halfMistakes, err)
				masterWordList = masterWordList[i+1:]
				testWordList = testWordList[j+1:]
				continue
			}

			err := error.NewTypingError(error.IncompleteWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)

			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		} else if spelling.IsSpellingMistake(masterWord, testWord) { // Check for spelling errors using Levenshtein distance.
			err := error.NewTypingError(error.SpellingMistake, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)

			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		} else if strings.EqualFold(nextMasterWord, nextTestWord) {
			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			err := error.NewTypingError(error.IncorrectWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			continue
		}

		// sliding window - find next match
		var halfErrorsNew, fullErrorsNew []error.TypingError
		masterWordList, testWordList, halfErrorsNew, fullErrorsNew = findNextMatch(masterWordList, testWordList, 3, masterWord, testWord, WordIndexInMaster)
		halfMistakes = append(halfMistakes, halfErrorsNew...)
		fullMistakes = append(fullMistakes, fullErrorsNew...)

		fmt.Println("-----------")
	}

	return fullMistakes, halfMistakes
}

func findNextMatch(masterWordList, testWordList []string, windowSize int, masterWord, testWord string, wordIndexInMaster int) ([]string, []string, []error.TypingError, []error.TypingError) {
	var fullMistakes, halfMistakes []error.TypingError
	findThese, ok := findNextWords(0, masterWordList, windowSize)
	if !ok {
		return masterWordList, testWordList, halfMistakes, fullMistakes
	}

	if len(findThese) <= 2 {
		masterWordList = masterWordList[1:]
		testWordList = testWordList[1:]
		err := error.NewTypingError(error.IncorrectWord, masterWord, testWord, wordIndexInMaster)
		fullMistakes = append(fullMistakes, err)
	}

	// find words in testWordList
	for i, tw := range testWordList {
		if !tokenize.IsAlphanumeric(tw) {
			err := error.NewTypingError(error.ExtraPunctuation, masterWord, tw, wordIndexInMaster)
			halfMistakes = append(halfMistakes, err)
		} else if testWords, ok := findNextWords(i, testWordList, windowSize); ok {
			if sameWords(findThese, testWords) {
				testWordList = testWordList[i:]
				break
			}
		}

		err := error.NewTypingError(error.IncorrectWord, masterWord, tw, wordIndexInMaster)
		fullMistakes = append(fullMistakes, err)

	}

	return masterWordList, testWordList, halfMistakes, fullMistakes
}

func sameWords(findThese, testWords []string) bool {
	for i, word := range findThese {
		if !(len(testWords) > i && strings.EqualFold(word, testWords[i])) {
			return false
		}
	}

	return true
}

func findNextWords(i int, testWordList []string, windowSize int) ([]string, bool) {
	var result []string
	counter := windowSize

	for j := i; j < len(testWordList); j++ {
		if len(result) == windowSize {
			return result, true
		}

		if (j + counter) > len(testWordList) { // insufficient wods left for windowSize
			return []string{}, false
		}

		tw := testWordList[j]
		if tokenize.IsAlphanumeric(tw) {
			result = append(result, tw)
			counter--
		}
	}

	return []string{}, false

}

func isSamePunctuation(w1, w2 string) bool {
	if w1 == "\n" {
		w1 = " "
	}
	if w2 == "\n" {
		w1 = " "
	}

	if !tokenize.IsAlphanumeric(w1) && !tokenize.IsAlphanumeric(w2) && w1 == w2 {
		return true
	}

	return false
}
