package main

type MinStack struct {
	stack    *[]int
	minStack *[]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack []int
	var minStack []int
	return MinStack{stack: &stack, minStack: &minStack}
}

func (this *MinStack) Push(x int) {
	minStackLen := len(*this.minStack)
	if minStackLen == 0 || x <= (*this.minStack)[minStackLen-1] {
		*this.minStack = append(*this.minStack, x)
	}
	*this.stack = append(*this.stack, x)
}

func (this *MinStack) Pop() {
	stackLen := len(*this.stack)
	minStackLen := len(*this.minStack)
	if stackLen == 0 {
		return
	}
	topInt := (*this.stack)[stackLen-1]
	*this.stack = (*this.stack)[:stackLen-1]
	if (*this.minStack)[minStackLen-1] == topInt {
		*this.minStack = (*this.minStack)[:minStackLen-1]
	}
}

func (this *MinStack) Top() int {
	stackLen := len(*this.stack)
	return (*this.stack)[stackLen-1]
}

func (this *MinStack) GetMin() int {
	minStackLen := len(*this.minStack)
	return (*this.minStack)[minStackLen-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
