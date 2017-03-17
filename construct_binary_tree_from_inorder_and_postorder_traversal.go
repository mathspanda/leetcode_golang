package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	rootIndex := index(inorder, rootVal)

	// partition point
	partitionIndex := len(inorder[:rootIndex])

	if rootIndex <= len(inorder) && partitionIndex <= len(postorder) {
		root.Left = buildTree(inorder[:rootIndex], postorder[:partitionIndex])
	}
	if rootIndex+1 < len(inorder) && partitionIndex < len(postorder) {
		root.Right = buildTree(inorder[rootIndex+1:], postorder[partitionIndex:len(postorder)-1])
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
