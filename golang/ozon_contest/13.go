package ozoncontest

import (
	"fmt"
)

func Task13Handler() bool {
	var text string
	fmt.Scan(&text)
	res := false
	var (
		isStart  = false
		isRe     = false
		isCancel = false
	)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'M':
			if isStart && !isCancel {
				return false
			}
			isCancel = false
			res = false
			isStart = true
		case 'R':
			if !isStart || isRe || isCancel {
				return false
			}
			isRe = true
		case 'C':
			if isCancel || !isStart {
				return false
			}
			isCancel = true
			isRe = false
		case 'D':
			if !isStart {
				return false
			}
			res = true
			isRe = false
			isStart = false
			isCancel = false
		}
	}

	return res
}

func Task13() {
	var count int
	fmt.Scan(&count)
	results := make([]bool, count)

	for i := 0; i < count; i++ {
		results[i] = Task13Handler()
	}

	for i := 0; i < count; i++ {
		if results[i] {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
