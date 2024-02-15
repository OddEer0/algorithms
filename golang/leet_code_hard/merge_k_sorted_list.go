package leetcodehard

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMinIndex(lists []*ListNode) (int, bool) {
	min := math.MaxInt
	found := false
	for i, node := range lists {
		if node != nil && !found {
			found = true
			min = i
		} else if node != nil && node.Val < lists[min].Val {
			min = i
		}
	}
	return min, found
}

func reduceMergeSortedList(head *ListNode, prev *ListNode, lists []*ListNode) *ListNode {
	if head == nil {
		ix, found := getMinIndex(lists)
		if !found {
			return nil
		}
		newHead := &ListNode{lists[ix].Val, nil}
		lists[ix] = lists[ix].Next
		return reduceMergeSortedList(newHead, newHead, lists)
	}
	ix, found := getMinIndex(lists)
	if !found {
		return head
	}
	newNode := &ListNode{lists[ix].Val, nil}
	prev.Next = newNode
	lists[ix] = lists[ix].Next
	return reduceMergeSortedList(head, newNode, lists)
}

func MergeKSortedList(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	return reduceMergeSortedList(nil, nil, lists)
}
