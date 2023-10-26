package main

import (
	"fmt"
	"time"
)

//Остановка горутины с помощью функции отмены
func main() {
	//Создаем канал и функцию отмены
	ch, cancel := count()

	//cпустя 2 секунды вызываем функцию отмены
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	//Читаем данные из канала, пока он не закроется
	for i := range ch {
		fmt.Println(i)
	}

}

func count() (<-chan int, func()) {
	//Канал ch, куда пишутся и откуда читаются данные
	ch := make(chan int)

	//Канал done при закрытии остановит горутину
	done := make(chan struct{})

	//Функция отмены закрывает канал done, что приводит к остановке горутины
	cancel := func() {
		close(done)
	}

	go func() {
		i := 0

		for {
			select {
			//Пока канал done закрыт, горутина остановится
			case <-done:
				close(ch)
				fmt.Println("goroutine stopped")
				return
			//Пока открыт канал done печатаем инкремент
			default:
				ch <- i
				i++
			}
		}
	}()
	return ch, cancel
}
