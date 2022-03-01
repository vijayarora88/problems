package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 5
	fmt.Println(countBits(n))
}

func countBits(n int) []int {
	arr := make([]int, 0)
	for i := 0; i <= n; i++ {
		formatInt := strconv.FormatInt(int64(i), 2)
		count := 0
		for _, value := range formatInt {
			if value%2 != 0 {
				count += 1
			}
		}
		arr = append(arr, count)
	}
	return arr
}
