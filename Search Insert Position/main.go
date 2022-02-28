package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s < e {
		m := (s + e) / 2
		if target > nums[m] {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if target <= nums[s] {
		return s
	}

	return s + 1
}
