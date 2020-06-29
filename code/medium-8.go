package code

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// StringRegex -
var StringRegex = regexp.MustCompile(`^(?i)(?P<sign>[+-]?)(?P<number>[0-9]+)[^0-9]?(.)*$`)

// MyAtoi -
func MyAtoi(str string) int {
	trimmedString := strings.TrimSpace(str)
	if trimmedString == "" || !isValid(trimmedString) {
		return 0
	}
	results := getParams(StringRegex, trimmedString)
	fmt.Println(results)
	sign, number := results["sign"], results["number"]
	fmt.Println("Sign = " + sign)
	var num int

	num, err := convertToNumber(number)

	// If I can't parse then I return either max or min value of Int 32
	if err != nil {
		if sign == "-" {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	// Since I am not passing sign
	if sign == "-" {
		num *= (-1)
	}

	if num >= math.MaxInt32 {
		num = math.MaxInt32
	}
	if num <= math.MinInt32 {
		num = math.MinInt32
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
func convertToNumber(str string) (number int, err error) {

	number, err = strconv.Atoi(str)
	if err != nil {

		return math.MaxInt32, err
	}
	return number, nil
}
