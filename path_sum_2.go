package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	result = [][]int{}
	if root == nil {
		return result
	}
	calculatePath(root, sum, []int{})
	return result
}

func calculatePath(root *TreeNode, sum int, temp []int) {
	if root == nil {
		if sum == 0 {
			result = append(result, append([]int{}, temp...))
		}
		return
	}
	if root.Left == nil {
		arg := append([]int{}, temp...)
		calculatePath(root.Right, sum-root.Val, append(arg, root.Val))
		return
	}
	if root.Right == nil {
		arg := append([]int{}, temp...)
		calculatePath(root.Left, sum-root.Val, append(arg, root.Val))
		return
	}
	arg := append([]int{}, temp...)
	calculatePath(root.Left, sum-root.Val, append(arg, root.Val))
	arg = append([]int{}, temp...)
	calculatePath(root.Right, sum-root.Val, append(arg, root.Val))
}
