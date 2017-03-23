package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	minLeftDepth := minDepth(root.Left)
	minRightDepth := minDepth(root.Right)
	if minLeftDepth < minRightDepth {
		return minLeftDepth + 1
	}
	return minRightDepth + 1
}
