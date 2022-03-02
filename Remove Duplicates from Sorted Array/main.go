package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	ptr1 := 0
	ptr2 := 0

	arr := make([]int, 0)

	for ptr2 < len(nums) {
		if ptr1 == ptr2 {
			arr = append(arr, nums[ptr2])
			ptr2++
			continue
		}

		if nums[ptr1] == nums[ptr2] {
			ptr2++
			continue
		}

		arr = append(arr, nums[ptr2])
		ptr1 = ptr2
		ptr2++
	}

	for i := 0; i < len(arr); i++ {
		nums[i] = arr[i]
	}

	return len(arr)
}
