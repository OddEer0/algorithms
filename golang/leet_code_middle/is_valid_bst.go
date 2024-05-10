package leetcodemiddle

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0, 10)
	cur := root
	prevVal := 0
	for cur != nil {
		stack = append(stack, cur)
		if cur.Left == nil {
			prevVal = cur.Val - 1
		}
		cur = cur.Left
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Val <= prevVal {
			return false
		}
		prevVal = current.Val
		if current.Right != nil {
			current = current.Right
			for current != nil {
				stack = append(stack, current)
				current = current.Left
			}
		}
	}

	return true
}
