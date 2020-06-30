package code

import (
	"math"
	"testing"

	"gitscm-sb.cisco.com/leet-code-problems/util"
)

func TestDivide(t *testing.T) {
	util.AssertEquals(t, Divide(10, 3), 3)
	util.AssertEquals(t, Divide(7, -3), -2)
	util.AssertEquals(t, Divide(math.MaxInt32, math.MaxInt32), 1)
	util.AssertEquals(t, Divide(math.MinInt32, math.MaxInt32), -1)
}
