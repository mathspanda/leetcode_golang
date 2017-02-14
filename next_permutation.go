package main

import (
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	i--

	j := len(nums) - 1
	for j > 0 && nums[j] <= nums[i] {
		j--
	}

	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
}
