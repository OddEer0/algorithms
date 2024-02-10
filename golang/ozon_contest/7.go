package ozoncontest

import (
	"fmt"
	"strconv"
	"strings"
)

func getWriteDocs() string {
	var totalPage int
	var writedPages string
	fmt.Scan(&totalPage)
	fmt.Scan(&writedPages)

	unwritePages := make([]bool, totalPage)

	w := strings.Split(writedPages, ",")
	for _, str := range w {
		nums := strings.Split(str, "-")
		if len(nums) == 2 {
			first, _ := strconv.Atoi(nums[0])
			second, _ := strconv.Atoi(nums[1])
			for k := first - 1; k < second; k++ {
				unwritePages[k] = true
			}
		} else {
			p, _ := strconv.Atoi(str)
			unwritePages[p-1] = true
		}
	}

	result := ""
	start := -1
	to := 0
	for i := 0; i < totalPage; i++ {
		if !unwritePages[i] {
			if start == -1 {
				start = i + 1
			}
			to++
		}
		if unwritePages[i] || i == totalPage-1 {
			if to == 1 {
				if result != "" {
					result += ","
				}
				if i == totalPage-1 && !unwritePages[i] {
					result += strconv.Itoa(i + 1)
				} else {
					result += strconv.Itoa(i)
				}
			} else if to > 1 {
				if result != "" {
					result += ","
				}
				if i == totalPage-1 && !unwritePages[i] {
					result += strconv.Itoa(start) + "-" + strconv.Itoa(i+1)
				} else {
					result += strconv.Itoa(start) + "-" + strconv.Itoa(i)
				}
			}
			to = 0
			start = -1
		}
	}

	return result
}

func Task7() {
	var count int
	fmt.Scan(&count)
	results := make([]string, count)
	for i := 0; i < count; i++ {
		results[i] = getWriteDocs()
	}

	for i := 0; i < count; i++ {
		fmt.Print(results[i])
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
