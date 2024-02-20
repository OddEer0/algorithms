package leetcodemiddle

func DivideTwoInt(dividend int, divisor int) int {
	var quotient, remainder int
	negFlag := false
	if dividend < 0 {
		dividend = -dividend
		negFlag = !negFlag
	}
	if divisor < 0 {
		divisor = -divisor
		negFlag = !negFlag
	}

	for i := 31; i >= 0; i-- {
		remainder = (remainder << 1) | ((dividend >> uint(i)) & 1)
		if remainder >= divisor {
			remainder -= divisor
			quotient |= (1 << uint(i))
		}
	}

	if negFlag {
		quotient = -quotient
		remainder = -remainder
	}

	if quotient >= 2147483648 {
		return 2147483647
	}

	return quotient
}
