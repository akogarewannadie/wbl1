package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5

	ch := make(chan int)

	go readFromChannel(ch)

	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	time.Sleep(time.Duration(N) * time.Second)
	fmt.Println("Программа завершена.")
}

func readFromChannel(ch <-chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println("Прочитано значение из канала:", val)
		}
	}
}
