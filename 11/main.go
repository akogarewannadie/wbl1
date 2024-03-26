package main

import "fmt"

func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	for k := range set1 {
		if set2[k] {
			result[k] = true
		}
	}

	return result
}

func main() {

	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	intersect := intersection(set1, set2)

	fmt.Println("Intersection:", intersect)
}
