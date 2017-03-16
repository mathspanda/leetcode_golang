package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := index(inorder, preorder[0])

	inorderLeft := inorder[:rootIndex]
	// partition point
	preorderIndex := len(inorderLeft) + 1
	if preorderIndex <= len(preorder) && rootIndex <= len(inorder) {
		root.Left = buildTree(preorder[1:preorderIndex], inorder[:rootIndex])
	}
	if preorderIndex < len(preorder) && rootIndex+1 < len(inorder) {
		root.Right = buildTree(preorder[preorderIndex:], inorder[rootIndex+1:])
	}
	return root
}

func index(ss []int, s int) int {
	for index, temp := range ss {
		if temp == s {
			return index
		}
	}
	return -1
}
