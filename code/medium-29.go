package code

import "math"

// Divide -
func Divide(dividend int, divisor int) int {

	sign := 1
	if (divisor < 0) != (dividend < 0) {
		sign = -1
	}
	var quo int

	// I'll work on only positive values and add sign in the end
	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		tmp := divisor
		powerOfTwo := 1
		for tmp*2 <= dividend {
			tmp *= 2
			powerOfTwo *= 2
		}
		quo += powerOfTwo
		dividend -= tmp
	}
	quo = handleOverFlow(quo, sign)
	return sign * quo
}

// Abs -
func abs(i int) int {
	if i < 0 {
		return (-1) * i
	}
	return i
}

func handleOverFlow(quo int, sign int) int {
	if quo > math.MaxInt32 {
		if sign == 1 {
			return math.MaxInt32
		}
		// Since I am handling sign in the end
		return (-1) * math.MinInt32

	}
	return quo
}
