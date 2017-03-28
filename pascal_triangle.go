package main

func generate(numRows int) [][]int {
	var result [][]int
	if numRows <= 0 {
		return result
	}
	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		var temp []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp = append(temp, 1)
				continue
			}
			temp = append(temp, result[i-1][j-1]+result[i-1][j])
		}
		result = append(result, temp)
	}
	return result
}
