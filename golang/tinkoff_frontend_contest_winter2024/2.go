package tinkoff_frontend_contest_winter2024

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func toInt(arr []string) ([]int, error) {
	result := make([]int, len(arr))
	var err error = nil

	for i, s := range arr {
		result[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("strconv error")
		}
	}

	return result, err
}

func sum(arr []int) int {
	result := 0

	for _, item := range arr {
		result += item
	}

	return result
}

func Task2() {
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	results := make([]string, count)

	for i := 0; i < count; i++ {
		scanner.Scan()
		_, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		arr, _ := toInt(s)
		summa := sum(arr)
		if summa%2 == 0 {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	for i := 0; i < count; i++ {
		fmt.Print(results[i])
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
