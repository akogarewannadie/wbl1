package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int)

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for res := range results {
		sum += res
	}

	fmt.Println("Сумма квадратов чисел:", sum)
}
