package code

import "sort"

// StrongPasswordChecker -
func StrongPasswordChecker(s string) int {
	hasDigit, hasLowerCase, hasUpperCase := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			hasDigit = 1
		}
		if s[i] <= 'z' && s[i] >= 'a' {
			hasLowerCase = 1
		}
		if s[i] <= 'Z' && s[i] >= 'A' {
			hasUpperCase = 1
		}
	}

	needChars := 3 - (hasDigit + hasLowerCase + hasUpperCase)
	if len(s) <= 4 {
		return 6 - len(s)
	}

	if len(s) == 5 {
		return max(6-len(s), needChars)
	}

	i := 0
	countThreeRepeat := 0
	for i < len(s) {
		if i < len(s) && i+1 < len(s) && i+2 < len(s) &&
			s[i] == s[i+1] && s[i] == s[i+2] {
			countThreeRepeat++
			i += 3
		} else {
			i++
		}
	}
	for len(s) >= 6 && len(s) <= 20 {
		return max(countThreeRepeat, needChars)
	}

	// now len(s) > 20
	needDelete := len(s) - 20
	result := 0
	if needChars > 0 {
		result += needChars
		countThreeRepeat -= needChars
	}

	if countThreeRepeat <= 0 || needDelete > 3*countThreeRepeat {
		result += needDelete
		return result
	}

	return mostComplicated(s, needChars)
}

func mostComplicated(s string, needChars int) int {
	needDelete := len(s) - 20
	repeat := []int{}
	idx := 0
	for idx < len(s) {
		j := idx
		for j < len(s) && s[idx] == s[j] {
			j++
		}
		count := j - idx
		if count > 2 {
			repeat = append(repeat, count)
		}
		idx = j
	}

	result := needChars
	for i := 0; i < needChars; i++ {
		sort.Ints(repeat)
		replaced := false
		for k := 2; k >= 0; k-- {
			for j := 0; j < len(repeat); j++ {
				if repeat[j] > 2 && repeat[j]%3 == k {
					repeat[j] -= 3
					replaced = true
				}
				if replaced {
					break
				}
			}
			if replaced {
				break
			}
		}
	}
	sort.Ints(repeat)
	for needDelete > 0 {
		br := true
		for k := 0; k <= 2; k++ {
			for j := 0; j < len(repeat); j++ {
				if repeat[j] > 2 && repeat[j]%3 == k && needDelete >= k+1 {
					repeat[j] -= (k + 1)
					needDelete -= (k + 1)
					result += (k + 1)
					br = false
				}
			}
		}
		if br {
			break
		}
	}

	result += needDelete
	for j := 0; j < len(repeat); j++ {
		result += repeat[j] / 3
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
