package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		i2 := i + nums[i]
		if i2 > reachable {
			reachable = i2
		}
	}
	return true
}
