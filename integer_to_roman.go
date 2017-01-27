package main

func intToRoman(num int) string {
	ints := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	result := ""
	for i := len(ints) - 1; i >= 0; i-- {
		for num >= ints[i] {
			result += romans[i]
			num -= ints[i]
		}
	}

	return result
}
