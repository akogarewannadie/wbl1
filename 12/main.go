package main

import "fmt"

func createSet(sequence []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range sequence {
		set[str] = true
	}

	return set
}

func main() {

	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(sequence)

	fmt.Println("Set:", set)
}
