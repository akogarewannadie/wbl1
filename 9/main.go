package main

import "fmt"

func main() {

	input := make(chan int)
	output := make(chan int)

	go processNumbers(input, output)

	go generateNumbers(input)

	for result := range output {
		fmt.Println(result)
	}
}

func generateNumbers(input chan<- int) {
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		input <- num
	}
	close(input)
}

func processNumbers(input <-chan int, output chan<- int) {
	for num := range input {
		result := num * 2
		output <- result
	}
	close(output)
}
