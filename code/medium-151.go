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

// ReverseWords2 -
func ReverseWords2(str string) string {
	var result string
	var word string

	var multipleSpaces bool
	for _, ele := range strings.TrimSpace(str) {

		if string(ele) == " " && multipleSpaces {
			continue
		} else if string(ele) == " " {
			multipleSpaces = true
			continue
		} else if string(ele) != " " && multipleSpaces {
			multipleSpaces = false
			result = word + " " + result
			word = ""
		}

		word += string(ele)

	}
	return strings.TrimSpace(word + " " + result)
}
