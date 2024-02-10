package leetcodemiddle

func FindFirstAndLastPositionOfElementInSortedArray(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		pivot := left + (right-left)/2
		switch {
		case nums[pivot] > target:
			right = pivot - 1
		case nums[pivot] < target:
			left = pivot + 1
		default:
			start, end := pivot, pivot

			for i := 1; start-1 >= 0 && nums[start-1] == nums[pivot]; i++ {
				start--
			}
			for i := 1; end+1 <= len(nums)-1 && nums[end+1] == nums[pivot]; i++ {
				end++
			}

			return []int{start, end}
		}
	}

	return []int{-1, -1}
}
