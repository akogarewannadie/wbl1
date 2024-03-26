package main

import (
	"fmt"
)

func reverseString(s string) string {

	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	str := "главрыба — абырвалг"
	reversedStr := reverseString(str)
	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевернутая строка:", reversedStr)
}
