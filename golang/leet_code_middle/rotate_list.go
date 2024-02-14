package leetcodemiddle

type ListNode struct {
	Val  int
	Next *ListNode
}

func reduceRotate(head *ListNode, curNode *ListNode, k, iter int) (int, *ListNode) {
	if curNode.Next == nil {
		newK := k % iter
		if newK == 0 {
			return -1, head
		}
		curNode.Next = head
		return newK - 1, curNode
	}
	newK, l := reduceRotate(head, curNode.Next, k, iter+1)
	if newK < 0 {
		return newK, l
	}
	if newK == 0 {
		curNode.Next = nil
		return -1, l
	}
	return newK - 1, curNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	_, newHead := reduceRotate(head, head, k, 1)
	return newHead
}
