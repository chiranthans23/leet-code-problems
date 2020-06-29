package util

import "testing"

// AssertEquals -
func AssertEquals(t *testing.T, actual, expected interface{}) {

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
