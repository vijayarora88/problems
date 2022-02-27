package main

import "fmt"

func main() {
	arr := []int{3, 3}
	fmt.Println(twoSum(arr, 6))
}

func twoSum(nums []int, target int) []int {
	arr := make(map[int]int)
	for index, value := range nums {
		if presentIndex, isPresent := arr[target-value]; isPresent {
			return []int{presentIndex, index}
		}
		arr[value] = index
	}

	return []int{}
}
