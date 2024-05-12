package my

import "math"

func MaxSumEvery2nd(sl []int) int {
	if len(sl) == 0 {
		return 0
	} else if len(sl) == 1 {
		return sl[0]
	}

	first, second := sl[0], sl[1]

	for i := 2; i < len(sl); i++ {
		max := math.Max(float64(second), float64(first+sl[i]))
		first = second
		second = int(max)
	}

	return second
}
