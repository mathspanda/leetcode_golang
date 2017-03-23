package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return has(root, sum)
}

func has(root *TreeNode, sum int) bool {
	if root == nil {
		if sum != 0 {
			return false
		} else {
			return true
		}
	}
	if root.Left == nil {
		return has(root.Right, sum-root.Val)
	}
	if root.Right == nil {
		return has(root.Left, sum-root.Val)
	}
	return has(root.Left, sum-root.Val) || has(root.Right, sum-root.Val)
}
