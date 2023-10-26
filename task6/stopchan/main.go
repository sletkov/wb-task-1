package main

import (
	"fmt"
	"time"
)

// Остановка горутины с помощью канала stop
func main() {
	//Создаем канал stop
	stop := make(chan bool)

	//Запускаем горутину
	go myProcess(stop)

	//Через 3 секунды записываем в канал stop значение true. Горутина должна остановиться
	time.Sleep(3 * time.Second)
	stop <- true

	fmt.Println("exit")
}

func myProcess(stop chan bool) {
	i := 0
	for {
		select {
		//Когда произойдет запись в канал true, горутина остановится
		case <-stop:
			fmt.Println("goroutine stopped")
			return
		//Пока не произошла запись в канал true, печатаем инкремент
		default:
			fmt.Println(i)
			i++
			time.Sleep(time.Second / 2)
		}
	}
}
