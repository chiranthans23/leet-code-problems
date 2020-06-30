package code

import "gitscm-sb.cisco.com/leet-code-problems/util"

// Divide -
func Divide(dividend int, divisor int) int {

	sign := 1
	if (divisor < 0) != (dividend < 0) {
		sign = -1
	}
	var quo int

	// I'll work on only positive values and add sign in the end
	dividend = util.Abs(dividend)
	divisor = util.Abs(divisor)

	for ; dividend >= divisor; quo++ {
		dividend -= divisor
	}

	return sign * quo
}
