package ozoncontest

import "fmt"

func task1() {
	var fieldCount int
	fmt.Scanf("%d", &fieldCount)
	results := make([]bool, fieldCount)

	for i := 0; i < fieldCount; i++ {
		arr := [4]int{0}
		result := true
		for k := 0; k < 10; k++ {
			var tmp int
			fmt.Scanf("%d", &tmp)
			arr[tmp-1]++
			if arr[tmp-1] > 5-tmp {
				result = false
			}
		}
		results[i] = result
	}

	for i, result := range results {
		if result {
			fmt.Print(success)
		} else {
			fmt.Print(fail)
		}
		if i != fieldCount-1 {
			fmt.Print("\n")
		}
	}
}
