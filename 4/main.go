package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	const numWorkers = 5

	dataChannel := make(chan int)
	done := make(chan struct{})

	startWorkers(numWorkers, dataChannel, done)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	close(done)
	fmt.Println("\nПрограмма завершена.")
}

func startWorkers(numWorkers int, dataChannel <-chan int, done <-chan struct{}) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerNum int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case data := <-dataChannel:
					fmt.Printf("Воркер %d: Получено значение: %d\n", workerNum, data)
				}
			}
		}(i + 1)
	}
	wg.Wait()
}
