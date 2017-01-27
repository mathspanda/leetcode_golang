package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	previous, current := head, head

	carryFlag := 0
	for l1 != nil && l2 != nil {
		current.Val = l1.Val + l2.Val + carryFlag
		carryFlag = current.Val / 10
		current.Val = current.Val % 10

		current.Next = &ListNode{}
		previous = current
		current = current.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		current.Val = l1.Val + carryFlag
		carryFlag = current.Val / 10
		current.Val = current.Val % 10

		current.Next = &ListNode{}
		previous = current
		current = current.Next
		l1 = l1.Next
	}
	for l2 != nil {
		current.Val = l2.Val + carryFlag
		carryFlag = current.Val / 10
		current.Val = current.Val % 10

		current.Next = &ListNode{}
		previous = current
		current = current.Next
		l2 = l2.Next
	}
	if carryFlag > 0 {
		current.Val = carryFlag
		current.Next = nil
	} else {
		previous.Next = nil
	}

	return head
}
