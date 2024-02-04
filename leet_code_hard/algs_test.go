package leetcodehard

import "testing"

func TestAlgs(t *testing.T) {
	cases := []struct {
		name   string
		nums1  []int
		nums2  []int
		result float64
	}{
		{
			name:   "case 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			result: 2.0,
		},
		{
			name:   "case 2",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			result: 2.5,
		},
		{
			name:   "case 3",
			nums1:  []int{1, 2, 3, 4},
			nums2:  []int{5, 6},
			result: 3.5,
		},
		{
			name:   "case 4",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4, 5, 6},
			result: 3.5,
		},
		// {
		// 	name:   "case 5",
		// 	nums1:  []int{2},
		// 	nums2:  []int{},
		// 	result: 2.0,
		// },
		// {
		// 	name:   "case 6",
		// 	nums1:  []int{},
		// 	nums2:  []int{2},
		// 	result: 2.0,
		// },
		{
			name:   "case 7",
			nums1:  []int{2, 3},
			nums2:  []int{},
			result: 2.5,
		},
		{
			name:   "case 8",
			nums1:  []int{},
			nums2:  []int{2, 3},
			result: 2.5,
		},
		{
			name:   "case 9",
			nums1:  []int{2, 3},
			nums2:  []int{-1, 4},
			result: 2.5,
		},
		{
			name:   "case 10",
			nums1:  []int{-1, 4},
			nums2:  []int{2, 3},
			result: 2.5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := FindMedianSortedArrays(tc.nums1, tc.nums2)
			if res != tc.result {
				t.Error("expected: ", tc.result, " but result: ", res)
			}
		})
	}
}
