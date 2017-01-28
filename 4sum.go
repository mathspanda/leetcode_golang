package main

import "sort"

// hash pair
func fourSum(nums []int, target int) [][]int {
	var results [][]int
	sort.Ints(nums)

	pairMap := make(map[int][][]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			tempSum := nums[i] + nums[j]
			pairMap[tempSum] = append(pairMap[tempSum], []int{i, j})
		}
	}

	for i := 0; i < len(nums)-3; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j >= i+2 && nums[j] == nums[j-1] {
				continue
			}
			if temp2Result, exist := pairMap[target-(nums[i]+nums[j])]; exist {
				firstAppend := true
				for k := 0; k < len(temp2Result); k++ {
					if j >= temp2Result[k][0] {
						continue
					}
					if firstAppend || nums[temp2Result[k][0]] != nums[temp2Result[k-1][0]] {
						firstAppend = false
						results = append(results, []int{nums[temp2Result[k][0]], nums[temp2Result[k][1]], nums[i], nums[j]})
					}
				}
			}
		}
	}

	return results
}
