package main

import (
	"fmt"
	"time"
)

//Остановка горутины с помощью канала done
func main() {
	//Создаем канал, куда будем записывать данные
	ch := make(chan int)

	//Создаем канал, сигнализурющий, что горутине нужно остановиться
	done := make(chan struct{})

	//Запускаем горутину
	go processData(ch, done)

	i := 0
	for {
		select {
		//Как только канал done закроется, горутина остановится
		case <-done:
			//Закрываем канал, куда пишем данные
			close(ch)
			fmt.Println("goroutine stopped")
			return
		//Пока канал done не закрылся, пишем в канал ch инкремент
		default:
			time.Sleep(500 * time.Millisecond)
			ch <- i
			i++
		}
	}

}

func processData(ch chan int, done chan struct{}) {
	//Спустя 4 секунды закрываем канал done
	go func() {
		time.Sleep(4 * time.Second)
		close(done)
	}()

	//Читаем данные из канала ch, пока канал done открыт
	for {
		select {
		case result := <-ch:
			fmt.Println(result)
		case <-done:
		}
	}
}
