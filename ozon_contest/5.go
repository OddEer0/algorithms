package ozoncontest

import "fmt"

func compression(seq []int) []int {
	result := make([]int, 0, len(seq))

	var (
		basicNum = 0
		seqCount = 0
		to       = 0
	)

	for i := 0; i < len(seq); i++ {
		isLast := i == len(seq)-1
		if i == 0 {
			basicNum = seq[i]
			result = append(result, basicNum)
		} else {
			if to == 0 {
				if seq[i] > basicNum {
					to = 1
				} else {
					to = -1
				}
			}
			if seq[i] == basicNum+to*(seqCount+1) {
				seqCount++
			} else {
				result = append(result, seqCount*to)
				seqCount = 0
				basicNum = seq[i]
				result = append(result, basicNum)
				to = 0
			}
		}
		if isLast {
			result = append(result, seqCount*to)
		}
	}

	return result
}

func task5() {
	var count int
	fmt.Scan(&count)
	results := make([][]int, count)

	for i := 0; i < count; i++ {
		var sequenceLength int
		fmt.Scan(&sequenceLength)
		sequence := make([]int, sequenceLength)

		for k := 0; k < sequenceLength; k++ {
			fmt.Scan(&sequence[k])
		}

		results[i] = compression(sequence)
	}

	for i, result := range results {
		fmt.Println(len(result))
		for k, num := range result {
			fmt.Print(num)
			if k != len(result)-1 {
				fmt.Print(" ")
			}
		}
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
