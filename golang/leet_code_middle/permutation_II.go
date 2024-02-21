package leetcodemiddle

import (
	"fmt"
	"strconv"
	"strings"
)

func sliceToSortedStr(sl []int) string {
	str := strings.Builder{}

	for i := range sl {
		str.WriteString(strconv.Itoa(sl[i]))
	}

	return str.String()
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func Permutation2(nums []int) [][]int {
	permCount := factorial(len(nums))
	result := make([][]int, 0, permCount)
	temp := make(map[string]bool, permCount)
	last := 1

	for i := 0; i < permCount; i++ {
		permute := make([]int, len(nums))
		if i == 0 {
			for j := 0; j < len(nums); j++ {
				permute[j] = nums[j]
			}
			key := sliceToSortedStr(permute)
			fmt.Println(key)
			if _, ok := temp[key]; !ok {
				result = append(result, permute)
				temp[key] = true
			} else {
				last++
			}
		} else {
			for j := 0; j < len(nums); j++ {
				permute[j] = result[i-last][j]
			}
			NextPermutation(permute)
			key := sliceToSortedStr(permute)
			fmt.Println(key)
			if _, ok := temp[key]; !ok {
				result = append(result, permute)
				temp[key] = true
			} else {
				last++
			}
		}
	}

	return result
}
