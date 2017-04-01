package main

func rangeBitwiseAnd(m int, n int) int {
	bit := 0
	for m != n {
		m = m / 2
		n = n / 2
		bit++
	}
	for i := 0; i < bit; i++ {
		m = m * 2
	}
	return m
}
