package main

import (
	"fmt"

	leetcodemiddle "github.com/OddEer0/algorithms/golang/leet_code_middle"
	"github.com/OddEer0/algorithms/golang/my"
)

func main() {
	str := "abcbada"
	uniqLen := my.LongestUniqueSubstring(str)
	fmt.Println(leetcodemiddle.UniquePaths(4, 4), uniqLen)
}
