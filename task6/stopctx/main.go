package main

import (
	"context"
	"fmt"
	"time"
)

//Остановка горутины с помощью context
func main() {
	//Создаем отменяемый контекст
	ctx, cancel := context.WithCancel(context.Background())

	//Вызываем горутину
	go myProcess(ctx)

	//Через 3 секунды отменяем контекст, в результате горутина остановится.
	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}

func myProcess(ctx context.Context) {
	i := 0
	for {
		select {
		//В случае отмены контекста сtx.Done() вернет канал, case отработает и горутина остановится
		case <-ctx.Done():
			fmt.Println("goroutine stopped")
			return
		//Пока контекст не отменен, печатаем инкремент
		default:
			fmt.Println(i)
			i++
		}
	}
}
