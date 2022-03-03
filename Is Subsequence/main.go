package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "axc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	end := 0
	arr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		s2 := string(s[i])
		for end < len(t) {
			t1 := string(t[end])
			if strings.EqualFold(s2, t1) {
				arr = append(arr, t1)
				end++
				break
			}
			end++
		}
	}

	return len(arr) == len(s)
}
