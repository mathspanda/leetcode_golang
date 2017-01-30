package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"

type Node struct {
	ele  *ListNode
	from int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].ele.Val < h[j].ele.Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := &ListNode{Val: 0}
	current := head
	nodes := NodeHeap{}
	for i := 0; i < len(lists); i++ {
		tempNode := lists[i]
		if tempNode == nil {
			continue
		}
		lists[i] = tempNode.Next
		nodes = append(nodes, Node{ele: tempNode, from: i})
	}
	if len(nodes) == 0 {
		return nil
	}
	heap.Init(&nodes)
	for len(nodes) > 0 {
		minNode := heap.Pop(&nodes).(Node)
		current.Next = minNode.ele
		current = current.Next
		if lists[minNode.from] != nil {
			tempNode := lists[minNode.from]
			lists[minNode.from] = tempNode.Next
			heap.Push(&nodes, Node{ele: tempNode, from: minNode.from})
		}
	}
	current.Next = nil
	return head.Next
}
