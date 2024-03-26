package main

import (
	"fmt"
)

func main() {
	var num int64 = 10
	fmt.Printf("Before setting: %d\n", num)

	i := 2
	num = setBit(num, i, 1)
	fmt.Printf("After setting bit %d to 1: %d\n", i, num)

	num = setBit(num, i, 0)
	fmt.Printf("After setting bit %d to 0: %d\n", i, num)
}

func setBit(num int64, i int, val int) int64 {

	mask := int64(^(1 << i))
	num &= mask

	if val == 1 {
		num |= (1 << i)
	}

	return num
}
