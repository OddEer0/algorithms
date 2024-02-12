package tinkoff_frontend_contest_winter2024

import (
	"fmt"
	"strconv"
	"strings"
)

func Task5() {
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	childrenCount, _ := strconv.Atoi(s[0])
	friendCount, _ := strconv.Atoi(s[1])
	eventCount, _ := strconv.Atoi(s[2])
	children := make([]int, childrenCount)
	friends := make([][]int, childrenCount)
	msg := make([]int, 0, 20)

	scanner.Scan()
	s = strings.Split(scanner.Text(), " ")
	for i := 0; i < childrenCount; i++ {
		children[i], _ = strconv.Atoi(s[i])
	}
	for i := 0; i < friendCount; i++ {
		scanner.Scan()
		s = strings.Split(scanner.Text(), " ")
		first, _ := strconv.Atoi(s[0])
		second, _ := strconv.Atoi(s[1])
		friends[first-1] = append(friends[first-1], second-1)
		friends[second-1] = append(friends[second-1], first-1)
	}

	for i := 0; i < eventCount; i++ {
		scanner.Scan()
		s = strings.Split(scanner.Text(), " ")
		switch s[0] {
		case "?":
			num, _ := strconv.Atoi(s[1])
			msg = append(msg, children[num-1])
		case "+":
			num, _ := strconv.Atoi(s[1])
			manul, _ := strconv.Atoi(s[2])
			friend := friends[num-1]
			for j := 0; j < len(friend); j++ {
				children[friend[j]] += manul
			}
		}
	}

	for i := 0; i < len(msg); i++ {
		fmt.Print(msg[i])
		if i != len(msg)-1 {
			fmt.Print("\n")
		}
	}
}
