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
	arg := []*leetcodehard.ListNode{
		&leetcodehard.ListNode{Val: 1, Next: &leetcodehard.ListNode{Val: 4, Next: &leetcodehard.ListNode{Val: 5, Next: nil}}},
		&leetcodehard.ListNode{Val: 1, Next: &leetcodehard.ListNode{Val: 3, Next: &leetcodehard.ListNode{Val: 4, Next: nil}}},
		&leetcodehard.ListNode{Val: 2, Next: &leetcodehard.ListNode{Val: 6, Next: nil}},
	}
	res := leetcodehard.MergeKSortedList(arg)
	PrintList(res)
}
