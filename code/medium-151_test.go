package code

import (
	"testing"

	"github.com/chiranthans23/leet-code-problems/util"
)

func TestReverseString(t *testing.T) {
	util.AssertEquals(t, ReverseWords("Am I"), "I Am")
	util.AssertEquals(t, ReverseWords(" Am I "), "I Am")
	util.AssertEquals(t, ReverseWords(" Am       I "), "I Am")
	util.AssertEquals(t, ReverseWords(" "), "")
	util.AssertEquals(t, ReverseWords("        I"), "I")
}
