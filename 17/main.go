package main

import (
	"fmt"
	"sort"
)

func binarySearch(slice []int, target int) int {
	index := sort.SearchInts(slice, target)
	if index < len(slice) && slice[index] == target {
		return index
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	index := binarySearch(slice, 5)
	if index != -1 {
		fmt.Printf("Число 5 найдено в позиции %d\n", index)
	} else {
		fmt.Println("Число 5 не найдено")
	}

	index = binarySearch(slice, 11)
	if index != -1 {
		fmt.Printf("Число 11 найдено в позиции %d\n", index)
	} else {
		fmt.Println("Число 11 не найдено")
	}
}
