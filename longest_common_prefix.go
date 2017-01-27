package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	index := 0
	for index < len(strs[0]) {
		tempRune := strs[0][index]
		flag := true
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || strs[i][index] != tempRune {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		index++
	}
	return strs[0][:index]
}
