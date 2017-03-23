package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	midNode := mid(head)
	root := new(TreeNode)
	root.Val = midNode.Val
	if head != midNode {
		root.Left = sortedListToBST(head)
	}
	if midNode != nil {
		root.Right = sortedListToBST(midNode.Next)
	}
	return root
}

func mid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preslow, slow, fast := head, head, head.Next
	for fast != nil && fast.Next != nil {
		preslow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if preslow != slow {
		preslow.Next = nil
	}
	return slow
}
