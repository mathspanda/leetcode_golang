package main

const (
	MaxInt = 0x7FFFFFFF
	MinInt = -MaxInt - 1
)

func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}

	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}

	if temp < MinInt || temp > MaxInt {
		return 0
	}

	return flag * temp
}
