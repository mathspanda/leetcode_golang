package main

func maxSubArray(nums []int) int {
	sum, max := 0, -2147483647
	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
