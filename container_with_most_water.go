package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	low, high := 0, len(height)-1
	max := 0
	for low < high {
		tempArea := min(height[low], height[high]) * (high - low)
		if tempArea > max {
			max = tempArea
		}
		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return max
}
