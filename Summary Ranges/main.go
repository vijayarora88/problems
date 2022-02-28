package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	if len(nums) <= 0 {
		return []string{}
	}

	result, start := []string{}, nums[0]

	for i := range nums {
		if i == len(nums)-1 || nums[i+1]-nums[i] > 1 {
			output := strconv.Itoa(start)
			if start < nums[i] {
				output += fmt.Sprintf("->%d", nums[i])
			}
			result = append(result, output)
			if i < len(nums)-1 {
				start = nums[i+1]
			}
		}
	}

	return result
}
