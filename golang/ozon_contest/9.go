package ozoncontest

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func analyzeGameFieldInput() [][]string {
	scan.Scan()
	s := strings.Split(scan.Text(), " ")
	row, _ := strconv.Atoi(s[0])
	field := make([][]string, row)
	for i := 0; i < row; i++ {
		scan.Scan()
		field[i] = strings.Split(scan.Text(), "")
	}

	return field
}

func isFrame(field [][]string, rowStart, colStart, rowEnd, colEnd int) (bool, int, int) {
	var (
		row = 0
		col = 0
	)
	for i := rowStart; i <= rowEnd; i++ {
		if field[i][colStart] == "*" {
			row++
			if i != rowStart && field[i][colStart+1] == "*" {
				break
			}
		} else {
			break
		}
	}
	if row < 3 {
		return false, 0, 0
	}
	for i := colStart; i <= colEnd; i++ {
		if field[rowStart][i] == "*" {
			col++
			if i != colStart && field[rowStart+1][i] == "*" {
				break
			}
		} else {
			break
		}
	}
	if col < 3 {
		return false, 0, 0
	}
	for i := rowStart; i < row; i++ {
		if field[i][colStart+col-1] != "*" {
			return false, 0, 0
		}
	}
	for i := colStart; i < col; i++ {
		if field[rowStart+row-1][i] != "*" {
			return false, 0, 0
		}
	}

	for i := rowStart; i < rowStart+row; i++ {
		field[i][colStart] = strconv.Itoa(col)
	}

	return true, row, col
}

func reduceFind(result *[]int, field [][]string, high, rowStart, colStart, rowEnd, colEnd int) {
	for i := rowStart; i <= rowEnd-2; i++ {
		for j := colStart; j <= colEnd-2; j++ {
			switch field[i][j] {
			case ".":
				continue
			case "*":
				ok, rowLen, colLen := isFrame(field, i, j, rowEnd, colEnd)
				if ok {
					*result = append(*result, high)
					if rowLen >= 7 && colLen >= 7 {
						reduceFind(result, field, high+1, i+2, j+2, i+rowLen-2, j+colLen-2)
					}
					j += colLen
				}
			default:
				step, _ := strconv.Atoi(field[i][j])
				j += step
			}
		}
	}
}

func analyzeGameField(field [][]string) []int {
	row, col := len(field), len(field[0])
	res := make([]int, 0, 10)
	reduceFind(&res, field, 0, 0, 0, row-1, col-1)
	return res
}

func Task9() {
	scan.Scan()
	count, _ := strconv.Atoi(scan.Text())
	results := make([][]int, count)

	for i := 0; i < count; i++ {
		field := analyzeGameFieldInput()
		results[i] = analyzeGameField(field)
		sort.Ints(results[i])
	}

	for i := 0; i < count; i++ {
		for j := 0; j < len(results[i]); j++ {
			fmt.Print(results[i][j], " ")
		}
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
