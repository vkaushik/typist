package spelling

import "github.com/vkaushik/typist/pkg/levenshtein"

// IsSpellingMistake checks if the given word is a spelling mistake based on Levenshtein distance.
func IsSpellingMistake(expected, actual string) bool {
	distance := levenshtein.LevenshteinDistance(expected, actual)
	maxLen := float64(max(len(expected), len(actual)))

	return float64(distance) <= maxLen/2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
