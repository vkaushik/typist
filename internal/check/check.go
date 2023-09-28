package check

import (
	"fmt"
	"strings"
	"time"

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

		// now either user have missed the word 
		// or typed an incorrect word - possible mistakes

		if len(masterWord) < 2 { // no more words in master list
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
		} else if spelling.IsSpellingMistake(masterWord, testWord) { // Check for spelling errors using Levenshtein distance.
			err := error.NewTypingError(error.SpellingMistake, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)

			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		} else if strings.HasPrefix(masterWord, testWord) || strings.HasPrefix(testWord, masterWord) { // Check for incomplete words (half-typed).
			err := error.NewTypingError(error.IncompleteWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)

			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			fmt.Println(err.Error())
			continue
		} else { // wrong word
			masterWordList = masterWordList[1:]
			testWordList = testWordList[1:]
			err := error.NewTypingError(error.IncorrectWord, masterWord, testWord, WordIndexInMaster)
			fullMistakes = append(fullMistakes, err)
			continue
		}

		fmt.Println("something went wrong, fix it", masterWord, " - ", testWord, " : ", WordIndexInMaster)
		// fmt.Println()
		// fmt.Println(masterWordList)
		// fmt.Println(testWordList)
		fmt.Println("-----------")
		time.Sleep(time.Second * 2)
	}

	return fullMistakes, halfMistakes
}
