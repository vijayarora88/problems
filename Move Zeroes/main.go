package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	zPtr := 0
	nzPtr := 0

	for i := 0; i < len(nums); i++ {
		if nums[nzPtr] == 0 {
			nzPtr++
			continue
		}

		temp := nums[nzPtr]
		nums[nzPtr] = nums[zPtr]
		nums[zPtr] = temp

		nzPtr++
		zPtr++
	}
}
