package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	minAbs := 0x7FFFFFFF
	minTarget := 0
	for i := 0; i < len(nums)-2; i++ {
		twoTarget := twoSumClosest(nums, target-nums[i], i+1, len(nums)-1)
		tempAbs := int(math.Abs(float64(target - twoTarget - nums[i])))
		if tempAbs < minAbs {
			minAbs = tempAbs
			minTarget = twoTarget + nums[i]
		}
	}
	return minTarget
}

func twoSumClosest(nums []int, target, begin, end int) int {
	var minAbs = 0x7FFFFFFF
	var minTarget = 0
	for begin < end {
		tempSum := nums[begin] + nums[end]
		if tempSum < target {
			begin++
		} else if tempSum > target {
			end--
		} else {
			return target
		}
		tempAbs := int(math.Abs(float64(tempSum - target)))
		if tempAbs < minAbs {
			minAbs = tempAbs
			minTarget = tempSum
		}
	}
	return minTarget
}
