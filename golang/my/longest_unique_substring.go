package my

func LongestUniqueSubstring(str string) int {
	if str == "" && len(str) == 1 {
		return len(str)
	}
	uniqueMap := make(map[rune]bool, 28)
	runes := []rune(str)
	var result, start, i int
	for i < len(runes) {
		sym := runes[i]
		if _, ok := uniqueMap[sym]; !ok {
			uniqueMap[sym] = true
			length := i - start + 1
			if length > result {
				result = length
			}
			i++
		} else {
			_, ok := uniqueMap[sym]
			for ok {
				delete(uniqueMap, sym)
				start++
				_, ok = uniqueMap[sym]
			}
		}
	}

	return result
}
