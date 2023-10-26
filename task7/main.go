package main

import (
	"fmt"
	"sync"
)

func main() {
	// создаем WaitGroup
	var wg sync.WaitGroup
	// создаем map
	m := make(map[int]int)
	//создаем map
	mutex := sync.RWMutex{}

	//Вызовем в цикле 10 горутин, которые будут конкурентно записывать в map
	for i := 0; i < 10; i++ {
		//Увеличим кол-во ожидаемых горутин в WaiGroup
		wg.Add(1)
		go func(i int) {
			// в отложенном режиме в самом конце уменьшим кол-во ожидаемых горутин в WaitGroup
			defer wg.Done()

			//Закрываем RWMutex для записи в map
			mutex.Lock()

			//Записываем данные в map
			m[i] = i + 1

			//в отложенном режиме открываем RWMutex, чтобы можно было писать
			defer mutex.Unlock()
		}(i)
	}

	//Ждем, пока все горутины отработают
	wg.Wait()

	//Выводим map
	fmt.Println(m)
}
