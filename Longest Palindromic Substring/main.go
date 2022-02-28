package main

import (
	"fmt"
)

func main() {
	str := "cbbd"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start := 0
	end := 2

	stringResult := ""

	for end <= len(s) {
		str := s[start:end]
		rev := reverse(str)
		if str == rev {
			stringResult = str
			start++
			end++
		} else {
			end++
		}
	}

	return stringResult
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
