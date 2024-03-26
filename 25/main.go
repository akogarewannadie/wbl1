package main

import (
	"fmt"
	"time"
)

func sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Microsecond)
}
func main() {
	fmt.Println("Начало работы программы")
	sleep(2000)
	fmt.Println("Программа завершена")
}
