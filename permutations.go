package main

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var result [][]int
	for i, num := range nums {
		temp := make([]int, len(nums)-1)
		copy(temp, nums[:i])
		copy(temp[i:], nums[i+1:])
		for _, r := range permute(temp) {
			result = append(result, append([]int{num}, r...))
		}
	}
	return result
}
