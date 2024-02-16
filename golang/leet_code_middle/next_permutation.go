package leetcodemiddle

func reverseSliceByAB(sl []int, start, end int) {
	for start < end {
		sl[start], sl[end] = sl[end], sl[start]
		start++
		end--
	}
}

func NextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	pivot := length - 1
	for i := pivot; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pivot = i
			break
		}
		if i == 1 {
			reverseSliceByAB(nums, 0, length-1)
			return
		}
	}
	left := pivot - 1
	for i := pivot; i < length; i++ {
		if nums[i] <= nums[pivot] && nums[left] < nums[i] {
			pivot = i
		}
	}
	nums[left], nums[pivot] = nums[pivot], nums[left]
	reverseSliceByAB(nums, left+1, length-1)
}
