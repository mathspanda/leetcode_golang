package main

const (
	MaxInt = 0x7FFFFFFF
	MinInt = -MaxInt - 1
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return MaxInt
	}
	if dividend == MinInt && divisor == -1 {
		return MaxInt
	}
	negFlag := 1
	if (dividend^divisor)>>31 == -1 {
		negFlag = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	result := 0
	for dividend >= divisor {
		shiftNum := uint(0)
		for dividend >= divisor<<shiftNum {
			shiftNum++
		}
		result += 1 << (shiftNum - 1)
		dividend -= divisor * (1 << (shiftNum - 1))
	}
	return result * negFlag
}
