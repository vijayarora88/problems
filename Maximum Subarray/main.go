package main

import "fmt"

func main() {
	arr := []int{-10, 3 - 2, -4, -2}
	fmt.Println(getLargestSum(arr))
}

func getLargestSum(arr []int) int {
	maxSum := 0
	currentSum := 0
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}
