package error

import (
	"encoding/json"
	"fmt"
)

func NewTypingError(remark, expected, actual string, wordIndexInMaster int) TypingError {
	return TypingError{
		Remark:            remark,
		Expected:          expected,
		Actual:            actual,
		WordIndexInMaster: wordIndexInMaster,
	}
}

type TypingError struct {
	Remark            string `json:"remark"`
	Expected          string `json:"expected"`
	Actual            string `json:"actual"`
	WordIndexInMaster int    `json:"wordIndexInMaster"`
}

func (o *TypingError) Error() string {
	return fmt.Sprintf("error: %s, expected: %s, actual: %s, WordIndexInMaster: %d", o.Remark, o.Expected, o.Actual, o.WordIndexInMaster)
}

func (o *TypingError) JsonString() string {
	m, _ := json.Marshal(o)
	return string(m)
}

func (o *TypingError) ShortString() string {
	return fmt.Sprintf(`{"%s", "%s", %d}`, o.Expected, o.Actual, o.WordIndexInMaster)
}
