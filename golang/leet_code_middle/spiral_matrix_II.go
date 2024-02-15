package leetcodemiddle

func GenerateSpiralMatrix(n int) [][]int {
	result, iteCnt, value := make([][]int, n), (n+1)/2, 1
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < iteCnt; i++ {
		for j := i; j < n-i; j++ {
			result[i][j] = value
			value++
		}

		for j := i + 1; j < n-i-1; j++ {
			result[j][n-1-i] = value
			value++
		}

		if i < iteCnt-1 || n%2 == 0 {
			for j := n - 1 - i; j >= i; j-- {
				result[n-1-i][j] = value
				value++
			}

			for j := n - 2 - i; j > i; j-- {
				result[j][i] = value
				value++
			}
		}
	}

	return result
}
