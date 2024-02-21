package leetcodemiddle

func Permutation(nums []int) [][]int {
	permCount := factorial(len(nums))
	result := make([][]int, permCount)

	for i := 0; i < permCount; i++ {
		permute := make([]int, len(nums))
		if i == 0 {
			for j := 0; j < len(nums); j++ {
				permute[j] = nums[j]
			}
			result[i] = permute
		} else {
			for j := 0; j < len(nums); j++ {
				permute[j] = result[i-1][j]
			}
			NextPermutation(permute)
			result[i] = permute
		}
	}

	return result
}
