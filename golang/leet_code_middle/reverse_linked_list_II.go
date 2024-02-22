package leetcodemiddle

func reduceReverseBetween(head, current, leftL *ListNode, left, right, iter int) (*ListNode, int) {
	if iter == left {
		leftL = current
	}
	if iter <= right {
		node, i := reduceReverseBetween(head, current.Next, leftL, left, right, iter+1)
		if i > 0 {
			node.Val, current.Val = current.Val, node.Val
			return node.Next, i - 1
		}
		return nil, 0
	}
	return leftL, (right - left + 1) / 2
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	reduceReverseBetween(head, head, nil, left, right, 1)
	return head
}
