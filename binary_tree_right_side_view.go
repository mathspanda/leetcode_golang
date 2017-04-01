package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		result = append(result, queue[len(queue)-1].Val)
		var temp []*TreeNode
		for _, item := range queue {
			if item.Left != nil {
				temp = append(temp, item.Left)
			}
			if item.Right != nil {
				temp = append(temp, item.Right)
			}
		}
		queue = temp
	}

	return result
}
