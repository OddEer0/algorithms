package ozoncontest

import (
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x, y int
}

func Task14Input() ([][]byte, Coords, Coords) {
	scan.Scan()
	t := strings.Split(scan.Text(), " ")
	row, _ := strconv.Atoi(t[0])
	res := make([][]byte, row)
	aCoord := Coords{-1, -1}
	bCoord := Coords{-1, -1}
	for i := 0; i < row; i++ {
		scan.Scan()
		res[i] = []byte(scan.Text())
		if aCoord.x < 0 || bCoord.x < 0 {
			for j := 0; j < len(res[i]); j++ {
				if res[i][j] == 'A' || res[i][j] == 'B' {
					if aCoord.y == -1 {
						aCoord.y = j
						aCoord.x = i
					} else {
						bCoord.y = j
						bCoord.x = i
					}
				}
			}
		}
	}

	return res, aCoord, bCoord
}

func Task14Validate(field [][]byte, x, y int) bool {
	row, col := len(field), len(field[0])
	return x >= 0 && x < row && y >= 0 && y < col && field[x][y] == '.'
}

func Task14Handle(field [][]byte, start, end Coords, sym byte) bool {
	rows, cols := len(field), len(field[0])
	if rows == 0 || cols == 0 {
		return false
	}

	toX := 0
	if start.x > end.x || (start.x == end.x && start.y > end.y) {
		toX = -1
	} else {
		toX = 1
	}

	x, y := start.x, start.y
	for x != end.x {
		dx := x + toX
		if Task14Validate(field, dx, y) {
			x = dx
			field[x][y] = sym
		} else {
			y += toX
			if Task14Validate(field, x, y) {
				field[x][y] = sym
			} else {
				return false
			}
		}
	}
	for y != end.y {
		y += toX
		if Task14Validate(field, x, y) {
			field[x][y] = sym
		} else {
			return false
		}
	}

	return true
}

func Task14() {
	scan.Scan()
	count, _ := strconv.Atoi(scan.Text())
	results := make([][][]byte, count)

	for i := 0; i < count; i++ {
		field, a, b := Task14Input()
		results[i] = field
		if a.x > -1 && b.x > -1 {
			aEnd, bEnd := Coords{len(field) - 1, len(field[0]) - 1}, Coords{0, 0}
			aSym := field[a.x][a.y] + 32
			bSym := field[b.x][b.y] + 32
			if a.x < b.x || (a.x == b.x && a.y < b.y) {
				bEnd.x = aEnd.x
				bEnd.y = aEnd.y
				aEnd.x = 0
				aEnd.y = 0
			}
			Task14Handle(field, a, aEnd, aSym)
			Task14Handle(field, b, bEnd, bSym)
		}
	}

	for i := 0; i < count; i++ {
		for j := 0; j < len(results[i]); j++ {
			fmt.Print(string(results[i][j]))
			if j != len(results[i])-1 {
				fmt.Print("\n")
			}
		}
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
