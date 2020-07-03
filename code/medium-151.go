package code

import (
	"strings"
)

// ReverseWords - Reverses sentence
func ReverseWords(str string) string {

	setOfWords := strings.Split(strings.TrimSpace(str), " ")
	str = ""
	for i := len(setOfWords) - 1; i >= 0; i-- {
		if setOfWords[i] == "" {
			continue
		}
		if i == 0 {
			str += setOfWords[0]
			break
		}
		str += setOfWords[i] + " "
	}

	return str
}
