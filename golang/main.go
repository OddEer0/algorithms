package main

import (
	"fmt"

	leetcodemiddle "github.com/OddEer0/algorithms/golang/leet_code_middle"
)

func main() {
	res := []int{5, 4, 7, 5, 3, 2}
	leetcodemiddle.NextPermutation(res)
	fmt.Println(res)
}
