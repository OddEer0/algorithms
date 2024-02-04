package ozoncontest

import "fmt"

func isCorrectYear(year int) bool {
	return year%400 == 0 || year%100 != 0 && year%4 == 0
}

func task2() {
	var dateCount int
	fmt.Scan(&dateCount)
	results := make([]bool, dateCount)

	for i := 0; i < dateCount; i++ {
		var day, month, year int
		fmt.Scan(&day, &month, &year)
		if month >= 8 {
			month++
		}
		if month%2 == 0 && day == 31 {
			results[i] = false
		} else if month == 2 {
			if day > 29 {
				results[i] = false
			} else if day == 29 {
				results[i] = isCorrectYear(year)
			} else {
				results[i] = true
			}
		} else {
			results[i] = true
		}
	}

	for i, result := range results {
		if result {
			fmt.Print(success)
		} else {
			fmt.Print(fail)
		}
		if i != dateCount-1 {
			fmt.Print("\n")
		}
	}
}
