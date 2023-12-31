package main

import (
	"context"
	"flag"
	"fmt"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Объявляем переменную, где будем хранить кол-во воркеров
var (
	workersCount int
)

//В функции init объявляем флаг
func init() {
	flag.IntVar(&workersCount, "workers-count", 2, "count of workers")
}

func main() {
	//Парсим флаг
	flag.Parse()

	//Инициализируем канал
	//Так как в канал пишутся произвольные данные, канал имеет тип chan interface{}
	ch := make(chan interface{})

	//Создаем отменяемый контекст
	//В качестве способа завершения работы воркера выбран контекст, так как этот пакет удобно использовать.
	//Там уже есть функция отмены, и функция Done, которая возвращает канал в случае отмены контекста.
	//Использование такого способа позволяет писать меньше кода, ведь нам не нужно создавать дополнительные переменные
	ctx, cancel := context.WithCancel(context.Background())
	//В отложенном режиме вызываем функцию отмены
	defer cancel()

	//Запускаем необходимое количество воркеров
	for i := 1; i <= workersCount; i++ {
		go worker(ctx, i, ch)
	}

	//Пишем данные в канал
	write(ch)
}

//Функция write в бесконечном цикле записывает в канал инкремент
func write(ch chan<- interface{}) {
	i := 0
	for {
		ch <- i
		i++
	}
}

func worker(ctx context.Context, id int, ch <-chan interface{}) {
	for {
		select {
		//Если контекст отменен, воркер останавливается
		case <-ctx.Done():
			return
		//Если контекст не отменен, воркер выводит в stdout значение, прочитанное из канала
		default:
			fmt.Println("Worker", id, ":", <-ch)
		}
	}
}
