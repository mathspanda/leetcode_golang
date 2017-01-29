package main

func letterCombinations(digits string) []string {
	digitMap := make(map[int][]rune)
	digitMap[2] = []rune{'a', 'b', 'c'}
	digitMap[3] = []rune{'d', 'e', 'f'}
	digitMap[4] = []rune{'g', 'h', 'i'}
	digitMap[5] = []rune{'j', 'k', 'l'}
	digitMap[6] = []rune{'m', 'n', 'o'}
	digitMap[7] = []rune{'p', 'q', 'r', 's'}
	digitMap[8] = []rune{'t', 'u', 'v'}
	digitMap[9] = []rune{'w', 'x', 'y', 'z'}

	var results []string
	if len(digits) == 0 {
		return results
	}
	combine(digits, 0, digitMap, &results, "")
	return results
}

func combine(digits string, strIndex int, digitMap map[int][]rune, results *[]string, temp string) {
	if strIndex == len(digits) && temp != "" {
		*results = append(*results, temp)
		return
	}
	letters, exist := digitMap[int(digits[strIndex]-'0')]
	if exist {
		for i := 0; i < len(letters); i++ {
			tempStr := temp[:] + string(letters[i])
			combine(digits, strIndex+1, digitMap, results, tempStr)
		}
	}
}
