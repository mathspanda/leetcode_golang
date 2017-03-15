package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var nodes []*TreeNode
	nodes = append(nodes, root)
	direction := false
	for len(nodes) > 0 {
		direction = !direction
		var temp []int
		var tempNode []*TreeNode
		for _, node := range nodes {
			if node != nil {
				if direction {
					temp = append(temp, node.Val)
				} else {
					temp = append([]int{node.Val}, temp...)
				}
				if node.Left != nil {
					tempNode = append(tempNode, node.Left)
				}
				if node.Right != nil {
					tempNode = append(tempNode, node.Right)
				}
			}
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
		nodes = tempNode
	}
	return result
}
