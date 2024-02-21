package leetcodemiddle

func uniquePathsWithObstaclesReduce(grid [][]int, m, n int, res map[[2]int]int, x, y int) int {
	if grid[x][y] == 1 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 1 {
		return uniquePathsWithObstaclesReduce(grid, m, n-1, res, x, y+1)
	}
	if n == 1 {
		return uniquePathsWithObstaclesReduce(grid, m-1, n, res, x+1, y)
	}
	val, ok := res[[2]int{x, y}]
	if ok {
		return val
	}

	result := 0

	result += uniquePathsWithObstaclesReduce(grid, m-1, n, res, x+1, y)
	result += uniquePathsWithObstaclesReduce(grid, m, n-1, res, x, y+1)

	if m-1 > 1 && n-1 > 1 {
		res[[2]int{x, y}] = result
	}

	return result
}

func UniquePaths2(obstacleGrid [][]int) int {
	mem := make(map[[2]int]int, len(obstacleGrid)*len(obstacleGrid[0]))
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}

	return uniquePathsWithObstaclesReduce(obstacleGrid, len(obstacleGrid), len(obstacleGrid[0]), mem, 0, 0)
}
