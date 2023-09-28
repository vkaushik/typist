package spelling

import "github.com/vkaushik/typist/pkg/levenshtein"

// IsSpellingMistake checks if the given word is a spelling mistake based on Levenshtein distance.
func IsSpellingMistake(expected, actual string) bool {
	distance := levenshtein.LevenshteinDistance(expected, actual)

	// user typed a longer word and master word is just 2 or 1 char long, LD needs to be just 1 to be a spelling mistake
	// this is done as in once instance "to" was compared with "the" and it was taken as spelling mistake whereas it is a word miss.
	if len(actual) > len(expected) && len(expected) <= 2 && distance > 1 {
		return false
	}

	return distance <= 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
