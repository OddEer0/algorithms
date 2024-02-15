package leetcodemiddle

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func SortColors(nums []int) {
	redCount, blueCount := 0, 0
	for i := 0; i < len(nums)-blueCount; i++ {
		switch nums[i] {
		case 0:
			if redCount != i {
				swap(nums, redCount, i)
			}
			redCount++
		case 2:
			for nums[i] == 2 && len(nums)-1-blueCount > i {
				swap(nums, len(nums)-1-blueCount, i)
				blueCount++
			}
			if nums[i] == 0 {
				if redCount != i {
					swap(nums, redCount, i)
				}
				redCount++
			}
		}
	}
}
