package main

func combine(n int, k int) [][]int {
	return combineWithStart(n, k, 1)
}

func combineWithStart(n int, k int, startIndex int) [][]int {
	var result [][]int
	for i := startIndex; i <= n; i++ {
		if k == 1 {
			result = append(result, []int{i})
		} else {
			for _, r := range combineWithStart(n, k-1, i+1) {
				result = append(result, append([]int{i}, r...))
			}
		}
	}
	return result
}
