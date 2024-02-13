package leetcodemiddle

func toStrCounter(str string) [26]uint8 {
	var res [26]uint8

	for i := range str {
		res[str[i]-'a']++
	}

	return res
}

func GroupAnagrams(strs []string) [][]string {
	temp := make(map[[26]uint8][]string, len(strs))

	for i := range strs {
		ctr := toStrCounter(strs[i])
		temp[ctr] = append(temp[ctr], strs[i])
	}

	result := make([][]string, 0, len(temp))
	for _, t := range temp {
		result = append(result, t)
	}

	return result
}
