package tinkoff_frontend_contest_winter2024

import (
	"fmt"
	"strconv"
	"strings"
)

func arrIterator(arr []int, start, end int, callback func(int, int) int) {
	for i := start - 1; i < end; i++ {
		arr[i] = callback(arr[i], i)
	}
}

func inputTask6() ([]int, []string) {
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	event, _ := strconv.Atoi(s[1])
	scanner.Scan()
	s = strings.Split(scanner.Text(), " ")
	arr, _ := toInt(s)
	events := make([]string, event)

	for i := 0; i < event; i++ {
		scanner.Scan()
		events[i] = scanner.Text()
	}

	return arr, events
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Task6() {
	arr, events := inputTask6()
	results := make([]int, 0, len(events))

	for i := 0; i < len(events); i++ {
		s := strings.Split(events[i], " ")
		switch s[0] {
		case "+":
			start, _ := strconv.Atoi(s[1])
			end, _ := strconv.Atoi(s[2])
			add, _ := strconv.Atoi(s[3])
			arrIterator(arr, start, end, func(item, _ int) int {
				return item + add
			})
		case "?":
			start, _ := strconv.Atoi(s[1])
			end, _ := strconv.Atoi(s[2])
			k, _ := strconv.Atoi(s[3])
			b, _ := strconv.Atoi(s[4])
			max := 0
			arrIterator(arr, start, end, func(item, index int) int {
				m := min(item, k*(index+1)+b)
				if m > max {
					max = m
				}
				return item
			})
			results = append(results, max)
		}
	}

	for i := 0; i < len(results); i++ {
		fmt.Print(results[i])
		if i != len(results)-1 {
			fmt.Print("\n")
		}
	}
}
