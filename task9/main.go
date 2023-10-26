package main

import (
	"fmt"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chan1 := make(chan int)
	chan2 := make(chan int)

	// в цикле пройдемся по всем значениям массива
	for _, val := range arr {
		// вызовем горутину, для чтение элемента массива и записи в chan1
		go func(val int) {
			time.Sleep(500 * time.Millisecond) //немного замедлим скорость чтения для наглядности вывода
			chan1 <- val
		}(val)

		// вызовем горутину, для чтение из chan1 x и записи в chan2 x^2
		go func() {
			x := <-chan1
			chan2 <- x * x
		}()

		// Прочитаем содержимое chan2 и выведем в stdout
		fmt.Println(<-chan2)
	}

	// закроем каналы
	close(chan1)
	close(chan2)
}
