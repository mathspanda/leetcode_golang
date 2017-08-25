package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	headNode := &ListNode{Next: head}
	curNode := headNode
	for curNode != nil && curNode.Next != nil {
		for curNode.Next != nil && curNode.Next.Val == val {
			curNode.Next = curNode.Next.Next
		}
		curNode = curNode.Next
	}

	return headNode.Next
}
