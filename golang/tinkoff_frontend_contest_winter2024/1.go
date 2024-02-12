package tinkoff_frontend_contest_winter2024

import "fmt"

func Task1() {
	var count int
	fmt.Scan(&count)
	results := make([]string, count)
	for i := 0; i < count; i++ {
		var text string
		fmt.Scan(&text)
		maps := map[rune]int{'T': 1, 'I': 1, 'N': 1, 'K': 1, 'O': 1, 'F': 2}
		for _, s := range text {
			if maps[s] > 0 {
				maps[s]--
			} else {
				results[i] = "No"
				break
			}
		}
		if results[i] != "No" {
			results[i] = "Yes"
		}
	}

	for i := 0; i < count; i++ {
		fmt.Print(results[i])
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
