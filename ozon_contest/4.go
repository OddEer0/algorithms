package ozoncontest

import (
	"fmt"
	"strconv"
)

func getCond(q [][2]string) []int {
	result := make([]int, len(q))
	minTemp := 15
	maxTemp := 30

	for i, item := range q {
		temp, _ := strconv.Atoi(item[1])
		if item[0] == ">=" {
			if temp > minTemp {
				minTemp = temp
			}
		} else {
			if temp < maxTemp {
				maxTemp = temp
			}
		}
		if minTemp > maxTemp {
			result[i] = -1
		} else {
			result[i] = minTemp
		}
	}

	return result
}

func task4() {
	var count int
	fmt.Scan(&count)
	results := make([][]int, count)

	for i := 0; i < count; i++ {
		var emCount int
		fmt.Scan(&emCount)
		res := make([][2]string, emCount)
		for k := 0; k < emCount; k++ {
			fmt.Scanf("%s %s", &res[k][0], &res[k][1])
		}
		results[i] = getCond(res)
	}

	for _, result := range results {
		for _, resItem := range result {
			fmt.Println(resItem)
		}
		fmt.Print("\n")
	}
}
