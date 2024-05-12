package my

func Intersection(a, b []int) []int {
	result := []int{}
	mappedElem := make(map[int]int)

	for _, elem := range a {
		if _, ok := mappedElem[elem]; ok {
			mappedElem[elem]++
		} else {
			mappedElem[elem] = 1
		}
	}

	for _, elem := range b {
		if count, ok := mappedElem[elem]; ok && count > 0 {
			mappedElem[elem]--
			result = append(result, elem)
		}
	}

	return result
}
