package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		var temp []int
		var tempNode []*TreeNode
		for _, node := range nodes {
			if node != nil {
				temp = append(temp, node.Val)
				if node.Left != nil {
					tempNode = append(tempNode, node.Left)
				}
				if node.Right != nil {
					tempNode = append(tempNode, node.Right)
				}
			}
		}
		if len(temp) > 0 {
			result = append([][]int{temp}, result...)
		}
		nodes = tempNode
	}
	return result
}
