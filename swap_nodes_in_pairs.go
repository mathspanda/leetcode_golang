package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		slow.Val, fast.Val = fast.Val, slow.Val
		slow = fast.Next
		if slow == nil {
			break
		}
		fast = slow.Next
		if fast == nil {
			break
		}
	}
	return head
}
