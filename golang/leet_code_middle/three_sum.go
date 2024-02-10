package leetcodemiddle

// TODO - Сделать до конца

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target, ignore int) ([]int, bool) {
	if len(nums) < 2 {
		return nil, false
	}

	hashTable := map[int]int{}

	for index, value := range nums {
		if index != ignore {
			diff := target - value
			if numIndex, ok := hashTable[diff]; ok {
				return []int{numIndex, index}, true
			} else {
				hashTable[value] = index
			}
		}
	}

	return nil, false
}

func isEqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isUniqueSlice(s [][]int, target []int) bool {
	for _, i := range s {
		if isEqualSlice(i, target) {
			return false
		}
	}
	return true
}

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		j := 0
		k := i
		for {
			inds, ok := twoSum(nums[j:], -nums[i], k)
			if !ok {
				break
			}
			j += inds[1] + 1
			k = i - j
			inds[0] += (j - inds[1] - 1)
			inds[1] += (j - inds[1] - 1)
			fmt.Println(inds, "i: ", i, "nums: ", nums, j)
			if inds[0] != i && inds[1] != i {
				res := []int{nums[i], nums[inds[0]], nums[inds[1]]}
				sort.Ints(res)
				if isUniqueSlice(result, res) {
					result = append(result, res)
				}
			}
		}
	}

	return result
}
