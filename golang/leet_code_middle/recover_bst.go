package leetcodemiddle

func RecoverBst(root *TreeNode) {
	stack := make([]*TreeNode, 0, 10)
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}

	var small, big, prev *TreeNode
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > current.Val {
			small = current
			if big == nil {
				big = prev
			}
		}
		prev = current
		if current.Right != nil {
			current = current.Right
			for current != nil {
				stack = append(stack, current)
				current = current.Left
			}
		}
	}
	if small != nil && big != nil {
		small.Val, big.Val = big.Val, small.Val
	}
}
