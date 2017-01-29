package main

func generateParenthesis(n int) []string {
	var results []string
	generate(n, n, &results, "")
	return results
}

func generate(left, right int, results *[]string, tempStr string) {
	if left == 0 && right == 0 {
		*results = append(*results, tempStr)
		return
	}
	if left != 0 {
		temp := tempStr[:] + "("
		generate(left-1, right, results, temp)
	}
	if right != 0 && left < right {
		temp := tempStr[:] + ")"
		generate(left, right-1, results, temp)
	}
}
