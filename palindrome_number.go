func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    for i, j := 0, len(str) - 1; i <= j;  {
    	if str[i] != str[j] {
    		return false
    	}
    	i++
    	j--
    }
    return true
}