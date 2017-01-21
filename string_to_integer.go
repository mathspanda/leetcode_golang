
import "strings"

const (
	MaxInt = 0x7FFFFFFF
	MinInt = -MaxInt - 1
)

func myAtoi(str string) int {
	flag := 1
	sum := 0
	str = strings.Trim(str, " ")
	for index, s := range str {
		if index == 0 && s == '-' {
			flag = -1
			continue
		}
		if index == 0 && s == '+' {
			flag = 1
			continue
		}
		temp := int(s - '0')
		if temp >= 0 && temp <= 9 {
			if sum*10+temp < sum {
				if flag == 1 {
					return MaxInt
				} else {
					return MinInt
				}
			}
			sum = sum*10 + temp
		} else {
			break
		}
	}
	result := flag * sum
	if result < MinInt {
		return MinInt
	}
	if result > MaxInt {
		return MaxInt
	}
	return result
}
