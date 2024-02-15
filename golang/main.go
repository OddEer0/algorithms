package main

import (
	"fmt"

	leetcodehard "github.com/OddEer0/algorithms/golang/leet_code_hard"
)

func PrintList(head *leetcodehard.ListNode) {
	if head != nil {
		fmt.Print(head.Val, " ")
		PrintList(head.Next)
	}
}

func main() {
	res := &leetcodehard.ListNode{1, &leetcodehard.ListNode{2, &leetcodehard.ListNode{3, &leetcodehard.ListNode{4, &leetcodehard.ListNode{5, nil}}}}}
	leetcodehard.ReverseKGroup(res, 3)
	PrintList(res)
}
