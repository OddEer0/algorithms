package leetcodehard

func reduceReverseKGroup(start *ListNode, current *ListNode, iter, k int) (*ListNode, int) {
	if current == nil {
		return nil, 0
	}
	if iter%k == 0 {
		reduceReverseKGroup(current.Next, current.Next, 1, k)
		current.Val, start.Val = start.Val, current.Val
		return start.Next, k/2 - 1
	} else {
		t, i := reduceReverseKGroup(start, current.Next, iter+1, k)
		if t == nil {
			return nil, 0
		}
		if i > 0 {
			current.Val, t.Val = t.Val, current.Val
		}
		return t.Next, i - 1
	}
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}

	reduceReverseKGroup(head, head, 1, k)

	return head
}
