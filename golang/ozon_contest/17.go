package ozoncontest

import (
	"fmt"
	"strconv"
)

func Task17() {
	scan.Scan()
	n, _ := strconv.Atoi(scan.Text())
	employees := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		scan.Scan()
		employees[scan.Text()] = true
	}

	scan.Scan()
	m, _ := strconv.Atoi(scan.Text())
	results := make([]int, m)

	for i := 0; i < m; i++ {
		scan.Scan()
		nick := scan.Bytes()

		if _, ok := employees[string(nick)]; ok {
			results[i] = 1
		} else {
			for j := 0; j < len(nick)-1; j++ {
				nick[j], nick[j+1] = nick[j+1], nick[j]
				if _, ok := employees[string(nick)]; ok {
					results[i] = 1
					break
				}
				nick[j], nick[j+1] = nick[j+1], nick[j]
			}
		}

	}

	for i := 0; i < m; i++ {
		fmt.Print(results[i])
		if i != m-1 {
			fmt.Print("\n")
		}
	}
}
