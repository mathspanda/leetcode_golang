package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := depth(root)
	return balanced
}

func depth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftBalanced := depth(root.Left)
	rightDepth, rightBalanced := depth(root.Right)

	if !leftBalanced || !rightBalanced {
		return -1, false
	}
	if leftDepth == rightDepth {
		return leftDepth + 1, true
	}
	if leftDepth == rightDepth+1 {
		return leftDepth + 1, true
	}
	if leftDepth+1 == rightDepth {
		return rightDepth + 1, true
	}
	if leftDepth > rightDepth {
		return leftDepth + 1, false
	}
	return rightDepth + 1, false
}
