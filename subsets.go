package main

func subsets(nums []int) [][]int {
	bitset := make([]bool, len(nums))
	return subsetsWithCount(nums, bitset, len(nums), 0)
}

func subsetsWithCount(nums []int, bitset []bool, count, startIndex int) [][]int {
	if count == 0 {
		var tempResult []inst
		for i, bit := range bitset {
			if bit {
				tempResult = append(tempResult, nums[i])
			}
		}
		return [][]int{tempResult}
	}
	var result [][]int
	for i := startIndex; i < len(nums); i++ {
		bitset[i] = true
		result = append(result, subsetsWithCount(nums, bitset, count-1, i+1)...)
		bitset[i] = false
		result = append(result, subsetsWithCount(nums, bitset, count-1, i+1)...)
	}
	return result
}
