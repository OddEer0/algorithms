package main

import (
	"fmt"

	leetcodemiddle "github.com/OddEer0/algorithms/golang/leet_code_middle"
)

func main() {
	res := []int{2, 3, 6, 7}
	per := leetcodemiddle.CombinationSum(res, 7)
	fmt.Println(per)
}
