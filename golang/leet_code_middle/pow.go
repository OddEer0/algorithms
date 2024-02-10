package leetcodemiddle

func MyPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n == 1:
		return x
	case n < 0:
		return 1 / MyPow(x, -n)
	}

	if n%2 == 0 {
		result := MyPow(x, n/2)
		return result * result
	}

	return x * MyPow(x, n-1)
}
