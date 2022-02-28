package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	arr := make(map[int]int)
	for _, value := range nums {
		if _, isPresent := arr[value]; isPresent {
			return true
		}
		arr[value] = value
	}
	return false
}
