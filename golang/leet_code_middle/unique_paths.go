package leetcodemiddle

func uniquePathsReduce(m, n int, res map[[2]int]int, x, y int) int {
	if m == 1 || n == 1 {
		return 1
	}
	val, ok := res[[2]int{x, y}]
	if ok {
		return val
	}

	result := 0

	result += uniquePathsReduce(m-1, n, res, x+1, y)
	result += uniquePathsReduce(m, n-1, res, x, y+1)

	if m-1 > 1 && n-1 > 1 {
		res[[2]int{x, y}] = result
	}

	return result
}

func UniquePaths(m int, n int) int {
	mem := make(map[[2]int]int, m*n/2)

	return uniquePathsReduce(m, n, mem, 0, 0)
}
