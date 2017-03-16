package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var max int = 0
	maxDepthTraversal(root, 0, &max)
	return max
}

func maxDepthTraversal(root *TreeNode, depth int, max *int) {
	if root == nil {
		if depth > *max {
			*max = depth
		}
		return
	}
	maxDepthTraversal(root.Left, depth+1, max)
	maxDepthTraversal(root.Right, depth+1, max)
}
