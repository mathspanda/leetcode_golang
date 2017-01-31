package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[count] = nums[i]
			count++
			prev = nums[i]
		}
	}
	return count
}
