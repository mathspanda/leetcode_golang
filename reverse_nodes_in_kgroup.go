package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	step := 0
	prev, cur := head, head
	var last *ListNode
	for cur != nil {
		cur = cur.Next
		step++
		if step == k {
			step = 0
			tmpHead, tmpLast := reverse(prev, cur)
			if last == nil {
				head = tmpHead
				last = tmpLast
			} else {
				last.Next = tmpHead
				last = tmpLast
			}
			prev = cur
		}
	}
	return head
}

func reverse(begin *ListNode, end *ListNode) (head *ListNode, last *ListNode) {
	if begin == nil || begin.Next == nil {
		return begin, begin
	}
	prev, cur := begin, begin.Next
	prev.Next = end
	for cur != end {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev, begin
}
