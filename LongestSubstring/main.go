package main

import (
	"fmt"
)

func main() {

	str := "dvdf"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var result int
	charIndexMap := make(map[uint8]int)
	var start int
	for end := 0; end < n; end++ {
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			result = max(result, end-start)
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}
			start = duplicateIndex + 1
		}
		charIndexMap[s[end]] = end
	}
	result = max(result, n-start)
	return result
}

// max returns the max of num1 and num2
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
