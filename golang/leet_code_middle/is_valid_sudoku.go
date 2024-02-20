package leetcodemiddle

func IsValidSudoku(board [][]byte) bool {
	for i := range board {
		cache := make(map[byte]bool, 9)
		cache2 := make(map[byte]bool, 9)
		for j := range board[i] {
			if board[i][j] != '.' {
				if _, ok := cache[board[i][j]]; ok {
					return false
				}
				cache[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if _, ok := cache2[board[j][i]]; ok {
					return false
				}
				cache2[board[j][i]] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		cache := make(map[byte]bool, 9)
		col := i % 3
		row := 0
		if i > 0 {
			row = i / 3
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				m := board[j+row*3][k+col*3]
				if m != '.' {
					if _, ok := cache[m]; ok {
						return false
					}
					cache[m] = true
				}
			}
		}
	}

	return true
}
