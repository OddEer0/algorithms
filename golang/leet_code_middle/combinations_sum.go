package leetcodemiddle

func reduceCombinationSum(result *[][]int, candidates, current []int, target, iter, sum int) {
	if sum == target {
		*result = append(*result, append([]int{}, current...))
	}

	if iter >= len(candidates) || sum >= target {
		return
	}

	current = append(current, candidates[iter])
	reduceCombinationSum(result, candidates, current, target, iter, sum+candidates[iter])
	current = current[:len(current)-1]
	reduceCombinationSum(result, candidates, current, target, iter+1, sum)
}

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0, len(candidates))

	current := make([]int, 0, len(candidates))

	reduceCombinationSum(&result, candidates, current, target, 0, 0)

	return result
}
