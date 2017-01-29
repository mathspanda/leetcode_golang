package main

func isValid(s string) bool {
	var runes []rune
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			runes = append(runes, r)
		} else {
			if len(runes) == 0 {
				return false
			}
			tempRune := runes[len(runes)-1]
			runes = runes[:len(runes)-1]
			if r == ')' && tempRune != '(' {
				return false
			}
			if r == '}' && tempRune != '{' {
				return false
			}
			if r == ']' && tempRune != '[' {
				return false
			}
		}
	}
	if len(runes) != 0 {
		return false
	}
	return true
}
