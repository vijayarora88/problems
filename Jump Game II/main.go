package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	max := 0
	longestJump := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {
		max = int(math.Max(float64(max), float64(i+nums[i])))
		if i == longestJump {
			steps++
			longestJump = max
		}
	}
	return steps
}
