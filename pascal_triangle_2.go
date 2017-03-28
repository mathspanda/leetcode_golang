package main

func getRow(rowIndex int) []int {
	var result []int
	if rowIndex < 0 {
		return result
	}
	result = append(result, 1)
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				result[0] = 1
				continue
			}
			if j == i {
				result = append(result, 1)
				continue
			}
			result[j] = result[j-1] + result[j]
		}
	}
	return result
}
