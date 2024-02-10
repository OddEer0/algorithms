package ozoncontest

import (
	"fmt"
	"regexp"
	"strings"
)

var task3_pattern = regexp.MustCompile(`[a-zA-Z]\d{1,2}[a-zA-Z]{2}`)

func Task3() {
	var count int
	fmt.Scan(&count)
	result := make([]string, count)

	for i := 0; i < count; i++ {
		var str string
		fmt.Scan(&str)
		match := task3_pattern.FindAllString(str, -1)
		if len(strings.Join(match, "")) == len(str) {
			result[i] = strings.Join(match, " ")
		} else {
			result[i] = "-"
		}
	}

	for i, str := range result {
		fmt.Print(str)
		if i != count-1 {
			fmt.Print("\n")
		}
	}

}
