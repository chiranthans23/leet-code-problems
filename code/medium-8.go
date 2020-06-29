package code

import (
	"regexp"
	"strings"
)

// StringRegex -
var StringRegex = regexp.MustCompile(`^(?i)[?P<sign>+-]?(?P<number>\d*)\D*$`)

func MyAtoi(str string) int {
	trimmedString := strings.TrimSpace(str)
	if trimmedString == "" || !isValid(trimmedString) {
		return 0
	}
	results := getParams(StringRegex, trimmedString)
	sign, number := results["sign"], results["number"]
	var num int
	num = convertToNumber(number)
	if sign == "-" {
		num *= (-1)
	}
	return num
}
func isValid(str string) bool {

	return StringRegex.MatchString(str)
}

func getParams(regEx *regexp.Regexp, url string) (paramsMap map[string]string) {

	match := regEx.FindStringSubmatch(url)

	if match == nil {
		return
	}

	paramsMap = make(map[string]string)
	for i, name := range regEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
func convertToNumber(str string) int {
	n := len(str)
	var number int
	pow := 1
	for i := n; i >= 0; i-- {
		number += (number * pow)
		pow *= 10
	}
	return number
}
