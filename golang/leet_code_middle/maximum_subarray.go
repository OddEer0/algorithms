package leetcodemiddle

func MaxSubArray(nums []int) int {
	max := 0
	current := 0
	isStart := true

	for i := 0; i < len(nums); i++ {
		if current+nums[i] > current {
			isStart = false
			current += nums[i]
		} else if !isStart {
			isStart = true
			if current > max {
				max = current
			}
			current = 0
		}
	}

	return max
}
