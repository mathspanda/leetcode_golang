package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := make([]int, len(nums))
	if len(nums) > 0 {
		max[0] = nums[0]
	}
	if len(nums) > 1 {
		max[1] = maxInt(nums[0], nums[1])
	}
	if len(nums) > 2 {
		max[2] = maxInt(max[1], max[0]+nums[2])
	}
	for i := 3; i < len(nums); i++ {
		max[i] = maxInt(max[i-1], maxInt(max[i-2]+nums[i], max[i-3]+nums[i]))
	}
	return max[len(nums)-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
