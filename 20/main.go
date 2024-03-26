package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	words := strings.Fields(s)
	length := len(words)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	reversedStr := strings.Join(words, " ")

	return reversedStr
}

func main() {

	str := "snow dog sun"
	reversedStr := reverseWords(str)
	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевернутая строка:", reversedStr)
}
