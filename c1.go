package fuckleetcode

import "math"

func divide(a, b int) int {
	if a == math.MinInt32 && b < 0 {
		return math.MaxInt32
	}

	sign := 1

	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		sign = -1
	}

	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}

	res := 0
	for a <= b {
		a -= b
		res++
	}

	return res * sign
}