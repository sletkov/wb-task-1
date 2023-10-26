package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	fmt.Println("hello")
	sleep(5)
	fmt.Println("bye")
}

// Реализуем функцию sleep с помощью бесконечного цикла,
// который прервется по истечению какого-то кол-ва секунд, переданных как аргумент функции
func sleep(seconds int) {
	// Время cейчас
	now := time.Now()
	for {
		//Если с момента now прошло seconds или больше секунд - прерываем цикл
		if time.Since(now).Seconds() >= float64(seconds) {
			break
		}
	}
}
