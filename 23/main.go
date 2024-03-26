package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	index := 2

	if index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	} else {
		fmt.Println("Индекс выходит за пределы длины слайса")
	}

	fmt.Println(slice)
}
