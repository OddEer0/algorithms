package leetcodehard

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		if len(nums2)%2 == 1 {
			return float64(nums2[len(nums2)/2])
		}
		return float64(nums2[len(nums2)/2] + nums2[(len(nums2)/2)-1])
	} else if len(nums2) == 0 {
		if len(nums1)%2 == 1 {
			return float64(nums1[len(nums1)/2])
		}
		return float64(nums1[len(nums1)/1] + nums1[(len(nums1)/2)-1])
	}
	return 0.0
}
