package main

import (
	"fmt"
	"sync"
	"time"
)

//Остановка горутины с помощью переменной stop
func main() {
	//Создаем переменную stop
	var stop bool
	//Создаем wg
	var wg sync.WaitGroup

	//Будем ждать одну горутины
	wg.Add(1)

	//Запускаем горутину
	go myProcess(&stop, &wg)

	// Спустя 3 секунды присваивам stop значение true. Горутина должна остановиться
	time.Sleep(3 * time.Second)
	stop = true

	//Дожидаемся выполнения горутины
	wg.Wait()
	fmt.Println("exit")
}

func myProcess(stop *bool, wg *sync.WaitGroup) {
	//В отложенном режиме сообщаем о выполнении горутины
	defer wg.Done()

	i := 0
	for {
		//Если stop == true, цикл прервется и горутина остановится
		if *stop {
			fmt.Println("goroutine stopped")
			return
		}

		// Пока stop != true, печатаем инкремент
		fmt.Println(i)
		i++
		time.Sleep(time.Second / 2)
	}
}
