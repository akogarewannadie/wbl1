package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(str string) bool {
	str = strings.ToLower(str)
	uniqueChars := make(map[rune]bool)
	for _, char := range str {
		if uniqueChars[char] {
			return false
		}
		uniqueChars[char] = true
	}
	return true
}
func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(areAllCharactersUnique(str1))
	fmt.Println(areAllCharactersUnique(str2))
	fmt.Println(areAllCharactersUnique(str3))

}
