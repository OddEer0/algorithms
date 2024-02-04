package leetcodeeasy

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}

	for index, value := range nums {
		diff := target - value
		if numIndex, ok := hashTable[diff]; ok {
			return []int{numIndex, index}
		} else {
			hashTable[value] = index
		}
	}

	return []int{0, 0}
}
