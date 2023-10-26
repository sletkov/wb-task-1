package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

//Вариант с использованием mutex

//Объявляем структуру Counter, которая будет содержать count и mutex
type Counter struct {
	count int
	mutex sync.Mutex
}

func main() {
	//Cоздадим WaitGroup
	var wg sync.WaitGroup
	//Инициализируем переменную c, хранящую ссылку на структуру Counter
	c := &Counter{}

	//В цикле запустим 10 горутин, которые конкуренто будут увеличивать значение счетчика
	for i := 0; i < 10; i++ {
		//Увеличим количество ожидаемых горутин на 1
		wg.Add(1)
		go func() {
			//В отложенном режиме уменьшим кол-во ожидаемых горутин
			defer wg.Done()

			//Закроем mutex, пока будет производится увеличение счетчика
			c.mutex.Lock()

			//В отложенном режиме откроем mutex
			defer c.mutex.Unlock()

			//Увеличим счетчик на 1
			c.count++
		}()
	}

	//Ждем пока отработают все горутины
	wg.Wait()

	//Выведем значение сounter
	fmt.Println(c.count)
}