package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string]int)
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			data[fmt.Sprintf("key%d", i)] = i
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		for key, value := range data {
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}
		mu.Unlock()
	}()

	wg.Wait()
}
