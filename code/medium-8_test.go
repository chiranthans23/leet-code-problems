package code

import (
	"math"
	"strconv"
	"testing"

	"gitscm-sb.cisco.com/leet-code-problems/util"
)

func TestMyAtoi(t *testing.T) {
	util.AssertEquals(t, MyAtoi("-9223372036854775809"), math.MinInt32)
	util.AssertEquals(t, MyAtoi("-9223372036854775809"), math.MinInt32)
	util.AssertEquals(t, MyAtoi("-92233720368"), math.MinInt32)
	util.AssertEquals(t, MyAtoi("-9223372"), -9223372)
	util.AssertEquals(t, MyAtoi("My number is -9223372"), 0)
	util.AssertEquals(t, MyAtoi("-9.223372"), -9)
	util.AssertEquals(t, MyAtoi(strconv.Itoa(math.MaxInt32)), math.MaxInt32)
	util.AssertEquals(t, MyAtoi(strconv.Itoa(math.MinInt32)), math.MinInt32)
	util.AssertEquals(t, MyAtoi(strconv.Itoa(math.MaxInt32)+"Hello"), math.MaxInt32)
	util.AssertEquals(t, MyAtoi("Hello"+strconv.Itoa(math.MaxInt32)), 0)

}
