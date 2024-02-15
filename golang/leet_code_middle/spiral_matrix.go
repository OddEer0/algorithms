package leetcodemiddle

func SpiralMatrix(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	var (
		to      = 0
		iteCnt  = (len(matrix) + 1) / 2
		iteCnt2 = (len(matrix[0]) + 1) / 2
	)

	for i := 0; i < iteCnt && i < iteCnt2; i++ {
		for j := to; j < len(matrix[0])-to; j++ {
			result = append(result, matrix[to][j])
		}

		for j := to + 1; j < len(matrix)-to-1; j++ {
			result = append(result, matrix[j][len(matrix[0])-1-to])
		}

		if i < iteCnt-1 || len(matrix)%2 == 0 {
			for j := len(matrix[0]) - 1 - to; j >= to; j-- {
				result = append(result, matrix[len(matrix)-1-to][j])
			}
		}

		if i < iteCnt2-1 || len(matrix[0])%2 == 0 {
			for j := len(matrix) - 2 - to; j > to; j-- {
				result = append(result, matrix[j][to])
			}
		}

		to++
	}

	return result
}
