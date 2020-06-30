package util

import "testing"

// AssertEquals -
func AssertEquals(t *testing.T, actual, expected interface{}) {

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

// Abs -
func Abs(i int) int {
	if i < 0 {
		return (-1) * i
	}

	return i
}
