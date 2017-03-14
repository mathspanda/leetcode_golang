package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricSub(root.Left, root.Right)
}

func isSymmetricSub(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val == q.Val {
		if isSymmetricSub(p.Right, q.Left) && isSymmetricSub(p.Left, q.Right) {
			return true
		}
	}
	return false
}
