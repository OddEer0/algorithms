package leetcodemiddle

import "math"

func reverseInt(x int) int {
	result := 0

	for {
		if x == 0 {
			break
		}
		result = result*10 + (x % 10)
		x /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
