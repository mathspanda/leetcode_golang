package main

import "fmt"

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		var firstRune rune
		count := 1
		var tempStr string
		for index, r := range result {
			if index == 0 {
				firstRune = r
				continue
			}
			if r == firstRune {
				count++
				continue
			}
			tempStr += fmt.Sprintf("%d%c", count, firstRune)
			firstRune = r
			count = 1
		}
		tempStr += fmt.Sprintf("%d%c", count, firstRune)
		result = tempStr
	}
	return result
}
