package main

func largestRectangleArea(heights []int) int {
	var stack []int
	sum := 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		top := stack[len(stack)-1]
		if heights[top] <= heights[i] {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		var tmp int
		if len(stack) == 0 {
			tmp = heights[top] * i
		} else {
			tmp = heights[top] * (i - stack[len(stack)-1] - 1)
		}
		if tmp > sum {
			sum = tmp
		}
		i--
	}
	return sum
}
