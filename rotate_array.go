package main

func rotate(nums []int, k int) {
	indexMap := make(map[int]int)
	k = k % len(nums)
	for i := 0; i < len(nums); i++ {
		indexMap[(i+k)%len(nums)] = nums[i]
	}
	for index, value := range indexMap {
		nums[index] = value
	}
}
