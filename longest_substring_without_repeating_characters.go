func lengthOfLongestSubstring(s string) int {
    startIndex, maxLength := 0, 0
    charIndexMap := make(map[rune]int)

    for index, c := range s {
    	i, ok := charIndexMap[c]
    	if ok && startIndex <= i {
    	    startIndex = i + 1
    	} else {
    	    if index-startIndex+1 > maxLength {
    	        maxLength = index - startIndex + 1
    	    }
    	}
    	charIndexMap[c] = index
    }

    return maxLength
}