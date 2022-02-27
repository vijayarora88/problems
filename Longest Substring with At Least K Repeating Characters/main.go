package main

import (
	"fmt"
	"math"
)

func main() {
	str := "aaabb"
	k := 3
	fmt.Println(longestSubstring(str, k))
}

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) < k {
		return 0
	}

	charMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		s2 := s[i : i+1]
		if value, isPresent := charMap[s2]; isPresent {
			value += 1
			charMap[s2] = value
			continue
		}
		charMap[s2] = 1
	}

	max := 0
	start := 0
	isTrimmed := false
	for i := 0; i < len(s); i++ {
		s2 := s[i : i+1]
		if charMap[s2] < k {
			max = int(math.Max(float64(max), float64(longestSubstring(s[start:i], k))))
			start = i + 1
			isTrimmed = true
			break
		}
	}

	if !isTrimmed {
		max = int(math.Max(float64(max), float64(len(s))))
	} else {
		max = int(math.Max(float64(max), float64(longestSubstring(s[start:], k))))
	}

	return max
}
