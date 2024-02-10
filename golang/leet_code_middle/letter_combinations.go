package leetcodemiddle

var letterCombinationsMap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func reduceLetterCombinations(result []string, digits string, cur string, i, j int) int {
	syms := letterCombinationsMap[rune(digits[i])]

	for _, sym := range syms {
		if i < len(digits)-1 {
			j = reduceLetterCombinations(result, digits, cur+string(sym), i+1, j)
		} else {
			result[j] += cur + string(sym)
			j++
		}
	}

	return j
}

func resSizeLetterCombinations(digits string) int {
	res := 1
	for _, digit := range digits {
		res *= len(letterCombinationsMap[digit])
	}
	return res
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, resSizeLetterCombinations(digits))

	reduceLetterCombinations(res, digits, "", 0, 0)

	return res
}
