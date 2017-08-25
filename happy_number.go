package main

func isHappy(n int) bool {
	loop_dict := make(map[int]bool)
	loop_dict[n] = true

	for {
		sum, temp := 0, n
		for temp != 0 {
			digit := temp % 10
			sum += digit * digit
			temp = temp / 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := loop_dict[sum]; ok {
			return false
		}
		loop_dict[sum] = true
		n = sum
	}

	return false
}
