package leetcodemiddle

func isPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func LongestPalindrome(s string) string {
	for end := len(s); end > 0; end-- {
		for start := 0; start <= len(s)-end; start++ {
			if isPalindrome(s, start, start+end-1) {
				return s[start : start+end]
			}
		}
	}

	return ""
}
