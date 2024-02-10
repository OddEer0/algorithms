package leetcodemiddle

func IntToRoman(num int) string {
	result := ""
	for num >= 1000 {
		result += "M"
		num -= 1000
	}
	if num >= 900 {
		result += "CM"
		num -= 900
	}
	for num >= 500 {
		result += "D"
		num -= 500
	}
	if num >= 400 {
		result += "CD"
		num -= 400
	}
	for num >= 100 {
		result += "C"
		num -= 100
	}
	if num >= 90 {
		result += "XC"
		num -= 90
	}
	for num >= 50 {
		result += "L"
		num -= 50
	}
	if num >= 40 {
		result += "XL"
		num -= 40
	}
	for num >= 10 {
		result += "X"
		num -= 10
	}
	if num >= 9 {
		result += "IX"
		num -= 9
	}
	for num >= 5 {
		result += "V"
		num -= 5
	}
	if num == 4 {
		result += "IV"
		num -= 4
	}
	for num > 0 {
		result += "I"
		num--
	}
	return result
}
