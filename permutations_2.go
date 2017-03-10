package main

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var result [][]int
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		temp := make([]int, len(nums)-1)
		copy(temp, nums[:i])
		copy(temp[i:], nums[i+1:])
		for _, r := range permuteUnique(temp) {
			result = append(result, append([]int{num}, r...))
		}
	}
	return result
}
