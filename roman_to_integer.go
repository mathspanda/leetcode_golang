package main

func romanToInt(s string) int {
	ints := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	riMap := make(map[string]int)
	for i := 0; i < len(ints); i++ {
		riMap[romans[i]] = ints[i]
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if nums, exist := riMap[s[i:i+2]]; exist {
				result += nums
				i++
				continue
			}
		}
		if nums, exist := riMap[s[i:i+1]]; exist {
			result += nums
		}
	}

	return result
}
