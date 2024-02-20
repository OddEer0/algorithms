package leetcodemiddle

func reduceRemoveNthFromEnd(head *ListNode, prev *ListNode, current *ListNode, n int) (*ListNode, int) {
	if current == nil {
		return head, 1
	}
	_, k := reduceRemoveNthFromEnd(head, current, current.Next, n)
	if k < 0 {
		return head, k
	} else if k == n {
		if prev == nil {
			return head.Next, -1
		}
		prev.Next = current.Next
		current.Next = nil
		return head, -1
	}
	return head, k + 1
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil && n == 1 {
		return nil
	}
	res, _ := reduceRemoveNthFromEnd(head, nil, head, n)
	return res
}
