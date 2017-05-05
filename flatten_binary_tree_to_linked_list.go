package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root)
}

func dfs(root *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	head, tail := root, root
	var leftHead, leftTail *TreeNode
	var rightHead, rightTail *TreeNode
	if root.Left != nil {
		leftHead, leftTail = dfs(root.Left)
	}
	if root.Right != nil {
		rightHead, rightTail = dfs(root.Right)
	}
	if leftHead != nil {
		root.Right = leftHead
		leftTail.Right = rightHead
		if rightTail != nil {
			tail = rightTail
		} else {
			tail = leftTail
		}
	} else {
		tail = rightTail
	}
	root.Left = nil
	return head, tail
}
