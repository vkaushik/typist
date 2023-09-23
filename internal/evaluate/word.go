package evaluate

import (
	"strings"

	"github.com/vkaushik/typist/pkg/spelling"
)

// Define error constants.
const (
	NoError             = iota // 0
	CapitalizationError        // 1
	SpellingError              // 2
	IncompleteWordError        // 3
	OtherError                 // 4
)

func CompareWords(reference, test string) int {
	// Check if the words are the same.
	if reference == test {
		return NoError
	}

	// Check for capitalization errors.
	if strings.EqualFold(reference, test) {
		return CapitalizationError
	}

	// Check for spelling errors using Levenshtein distance.
	if spelling.IsSpellingMistake(reference, test) {
		return SpellingError
	}

	// Check for incomplete words (half-typed).
	if strings.HasPrefix(reference, test) || strings.HasPrefix(test, reference) {
		return IncompleteWordError
	}

	// If none of the above conditions are met.
	return OtherError
}

// // Function to check for spelling errors using Levenshtein distance.
// func isSpellingError(reference, test string) bool {
// 	const maxEditDistance = 2 // Define the maximum allowed Levenshtein distance.

// 	m, n := len(reference), len(test)
// 	if abs(m-n) > maxEditDistance {
// 		return true // Length difference is beyond the allowed threshold.
// 	}

// 	// Initialize a matrix to store Levenshtein distances.
// 	matrix := make([][]int, m+1)
// 	for i := range matrix {
// 		matrix[i] = make([]int, n+1)
// 	}

// 	// Initialize the first row and column of the matrix.
// 	for i := 0; i <= m; i++ {
// 		matrix[i][0] = i
// 	}
// 	for j := 0; j <= n; j++ {
// 		matrix[0][j] = j
// 	}

// 	// Calculate Levenshtein distances.
// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			cost := 0
// 			if reference[i-1] != test[j-1] {
// 				cost = 1
// 			}
// 			matrix[i][j] = min(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+cost)
// 		}
// 	}

// 	// The Levenshtein distance between the words is the value in the bottom-right cell of the matrix.
// 	return matrix[m][n] > maxEditDistance
// }

// // Helper function to calculate the minimum of three integers.
// func min(a, b, c int) int {
// 	if a <= b && a <= c {
// 		return a
// 	} else if b <= a && b <= c {
// 		return b
// 	}
// 	return c
// }

// // Helper function to calculate the absolute value of an integer.
// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

func GetErrorString(i int) string {
	switch i {
	case NoError:
		return "Both words are the same."
	case CapitalizationError:
		return "Capitalization error."
	case SpellingError:
		return "Spelling error."
	case IncompleteWordError:
		return "Incomplete word."
	case OtherError:
		return "unequal words error."
	}

	return "something went wrong"
}
