package main

import "fmt"

func main() {
	var num uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	count := uint32(0)
	for ; num > uint32(0); num >>= 1 {
		count += num & 1
	}
	return int(count)
}
