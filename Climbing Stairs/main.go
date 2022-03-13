package main

import "fmt"

func main() {
	n := 10
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	x := 1
	y := 2

	for i := 3; i <= n; i++ {
		temp := y
		y += x
		x = temp
	}

	return y
}
