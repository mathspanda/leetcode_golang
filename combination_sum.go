package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	combine(candidates, []int{}, 0, target, &result)
	return result
}

func combine(candidates, temps []int, startIndex, target int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		real := make([]int, len(temps))
		copy(real, temps)
		*result = append(*result, real)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		temps = append(temps, candidates[i])
		combine(candidates, temps, i, target-candidates[i], result)
		temps = temps[:len(temps)-1]
	}
	return
}
