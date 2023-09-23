package evaluate

type TypingError struct {
	Error string
	Expected string
	Actual string
}

func GetErrors(master, test string) ([]TypingError, []TypingError) {
	

	return nil, nil
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