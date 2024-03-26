package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func QuickSort(s []int) {
	sort.Sort(IntSlice(s))
}

func main() {

	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("Before sorting:", arr)

	QuickSort(arr)
	fmt.Println("After sorting:", arr)
}
