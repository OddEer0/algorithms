package leetcodehard

func getArrayMedian(nums []int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	mid := len(nums) / 2
	if len(nums)%2 == 1 {
		return float64(nums[mid])
	}
	return float64(nums[mid]+nums[mid-1]) / 2
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	switch {
	case length == 0:
		return 0
	case len(nums1) == 0:
		return getArrayMedian(nums2)
	case len(nums2) == 0:
		return getArrayMedian(nums1)
	case length == 2:
		return float64(nums1[0]+nums2[0]) / 2
	}

	// 2 array non empty len >= 3
	var (
		i           = 0
		j           = 0
		mid         = (length - 1) / 2
		isNotFirst  = false
		isNotSecond = false
	)

	for i+j < mid {
		if nums1[i] < nums2[j] {
			if i == len(nums1)-1 {
				isNotFirst = true
			}
			if isNotFirst {
				j++
			} else {
				i++
			}
		} else {
			if j == len(nums2)-1 {
				isNotSecond = true
			}
			if isNotSecond {
				i++
			} else {
				j++
			}
		}
	}

	switch {
	case length%2 == 1:
		switch {
		case isNotFirst:
			return float64(nums2[j-1])
		case isNotSecond:
			return float64(nums1[i-1])
		case nums1[i] < nums2[j]:
			return float64(nums1[i])
		default:
			return float64(nums2[j])
		}
	default:
		switch {
		case isNotFirst:
			return float64(nums2[j]+nums2[j-1]) / 2
		case isNotSecond:
			return float64(nums1[i]+nums1[i-1]) / 2
		case j != len(nums2)-1 && nums1[i] > nums2[j] && nums1[i] > nums2[j+1]:
			return float64(nums2[j]+nums2[j+1]) / 2
		case i != len(nums1)-1 && nums2[j] > nums1[i] && nums2[j] > nums1[i+1]:
			return float64(nums1[i]+nums1[i+1]) / 2
		default:
			return float64(nums1[i]+nums2[j]) / 2
		}
	}
}
