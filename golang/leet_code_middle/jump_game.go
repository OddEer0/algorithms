package leetcodemiddle

// На основе JumpGame II - да, я сперва 2 часть сделал. Лень было делать по другому))))
func JumpGame(nums []int) bool {
	last := len(nums) - 1
	if len(nums) == 1 {
		return true
	}
	if nums[0] >= last {
		return true
	}

	res := 0
	for i := 0; i < len(nums); {
		max := -1
		newIndex := 0
		for j := 1; j <= nums[i]; j++ {
			if nums[i+j] >= last-(i+j) {
				return true
			}
			length := nums[i+j] - nums[i] + j
			if length >= max {
				newIndex = i + j
				max = length
			}
		}
		if max < 1 {
			return false
		}
		i = newIndex
		res++
	}

	return false
}
