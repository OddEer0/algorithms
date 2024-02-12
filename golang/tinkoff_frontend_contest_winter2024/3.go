package tinkoff_frontend_contest_winter2024

import (
	"fmt"
	"strconv"
	"strings"
)

func Task3() {
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	deviceCount, _ := strconv.Atoi(s[0])
	credit, _ := strconv.Atoi(s[1])
	scanner.Scan()
	devices := strings.Split(scanner.Text(), " ")

	max := 0
	for i := 1; i <= credit; i++ {
		cur := i
		for j := 0; j < deviceCount; j++ {
			price, _ := strconv.Atoi(devices[j])
			if cur >= price {
				cur -= price
			}
		}
		if cur > max {
			max = cur
		}
	}

	fmt.Print(max)
}
