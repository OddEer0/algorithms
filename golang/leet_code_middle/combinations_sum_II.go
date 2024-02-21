package leetcodemiddle

import (
	"sort"
)

func reduceCombinationSum2(result *[][]int, candidates, current []int, target, iter int, temp map[string]bool) {
	if target == 0 {
		key := sliceToSortedStr(current)
		if _, ok := temp[key]; !ok {
			temp[key] = true
			temp := make([]int, len(current))
			copy(temp, current)
			*result = append(*result, temp)
		}
		return
	}

	for i := iter; i < len(candidates); i++ {
		if i > iter && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		current = append(current, candidates[i])
		reduceCombinationSum2(result, candidates, current, target-candidates[i], i+1, temp)
		current = current[:len(current)-1]
	}
}

func CombinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0, len(candidates))

	current := make([]int, 0, len(candidates))
	sort.Ints(candidates)
	temp := make(map[string]bool)

	reduceCombinationSum2(&result, candidates, current, target, 0, temp)

	return result
}
