package leetcodemiddle

import "math"

func myAtoi(s string) int {
	sign := 0
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && sign == 0 {
			continue
		}
		if s[i] == '-' && sign == 0 {
			sign = 1
		} else if s[i] == '+' && sign == 0 {
			sign = 2
		} else if s[i] <= '9' && s[i] >= '0' {
			if sign == 0 {
				sign = 2
			}
			result = result*10 + (int(s[i]) - int('0'))
			if result < math.MinInt32 || result > math.MaxInt32 {
				if sign == 1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	if sign == 1 {
		result = -result
	}

	return result
}
