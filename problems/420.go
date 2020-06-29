package code

import (
	"fmt"
	"strconv"
	"strings"
)

func strongPasswordChecker(s string) int {
	var numAdds int

	var numDeletes int
	var numLength int
	n := len(s)
	if n == 0 {
		return 6
	}

	if !checkForNumber(s) {
		numAdds++
	}
	if !checkForLowerCaseLetter(s) {
		numAdds++
	}
	if !checkForUpperCaseLetter(s) {
		numAdds++
	}

	triples := checkForTripleCharacters(s)
	numAdds = Max(numAdds, triples)

	if n < 6 {
		numLength = 6 - n
		numAdds = Max(numAdds, numLength)
	} else if n > 20 {
		numDeletes = n - 20
		numAdds = (numDeletes + numAdds)
	}

	return numAdds
}

func checkForTripleCharacters(s string) int {
	var triples int
	tripsCounter := 1

	curr := s[0]
	for i := 1; i < len(s)-2; i++ {
		if tripsCounter == 3 {
			tripsCounter = 1
			curr = s[i]
			continue
		}
		if s[i] == curr {
			tripsCounter++
			if tripsCounter >= 3 {
				triples++
			}
		}

	}
	fmt.Println(triples)
	return triples
}

func checkForNumber(s string) bool {

	for i := 0; i < 10; i++ {
		if strings.Contains(s, strconv.Itoa(i)) {
			return true
		}
	}
	return false

}
func checkForUpperCaseLetter(s string) bool {
	for i := 0; i < 26; i++ {
		if strings.Contains(s, toCharUpper(i)) {
			return true
		}
	}
	return false
}
func checkForLowerCaseLetter(s string) bool {
	for i := 0; i < 26; i++ {
		if strings.Contains(s, toCharLower(i)) {
			return true
		}
	}
	return false
}
func toCharUpper(i int) string {
	return string(rune('A' + i))
}
func toCharLower(i int) string {
	return string(rune('a' + i))
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
