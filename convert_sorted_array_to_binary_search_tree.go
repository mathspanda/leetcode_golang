package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIdx := len(nums) / 2
	root := new(TreeNode)
	root.Val = nums[midIdx]
	if midIdx <= len(nums) {
		root.Left = sortedArrayToBST(nums[:midIdx])
	}
	if midIdx+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[midIdx+1:])
	}
	return root
}
