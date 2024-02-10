package leetcodemiddle

func JumpGame2(nums []int) int {
	last := len(nums) - 1
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= last {
		return 1
	}

	res := 0
	for i := 0; i < len(nums); {
		max := -1
		newIndex := 0
		for j := 1; j <= nums[i]; j++ {
			if nums[i+j] >= last-(i+j) {
				return res + 2
			}
			length := nums[i+j] - nums[i] + j
			if length >= max {
				newIndex = i + j
				max = length
			}
		}
		i = newIndex
		res++
	}

	return 0
}
