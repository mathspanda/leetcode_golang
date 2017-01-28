package main

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] {
			continue
		}
		temp2Result := twoSum(nums, -nums[i], i+1, len(nums)-1)
		if len(temp2Result) != 0 {
			for j := 0; j < len(temp2Result); j++ {
				temp3Result := make([]int, 3)
				temp3Result[0] = nums[i]
				temp3Result[1] = temp2Result[j][0]
				temp3Result[2] = temp2Result[j][1]
				results = append(results, temp3Result)
			}
		}
	}

	return results
}

func twoSum(nums []int, target, begin, end int) [][]int {
	var results [][]int
	for begin < end {
		tempSum := nums[begin] + nums[end]
		if tempSum < target {
			begin++
		} else if tempSum > target {
			end--
		} else {
			tempResult := make([]int, 2)
			tempResult[0] = nums[begin]
			tempResult[1] = nums[end]
			results = append(results, tempResult)
			begin++
			end--
			// distinct
			for begin < end && nums[begin] == nums[begin-1] {
				begin++
			}
			for begin < end && nums[end] == nums[end+1] {
				end--
			}
		}
	}
	return results
}
