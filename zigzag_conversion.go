package main

func convert(s string, numRows int) string {
	result := make([]string, numRows)
	for i := 0; i < len(s); {
		for j := 0; j < numRows && i < len(s); j++ {
			result[j] += string(s[i])
			i++
		}
		for j := numRows - 2; j > 0 && i < len(s); j-- {
			result[j] += string(s[i])
			i++
		}
	}
	resultStr := ""
	for i := 0; i < len(result); i++ {
		resultStr += result[i]
	}
	return resultStr
}
