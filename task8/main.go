package main

import (
	"fmt"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	//Объявляем переменную типа int64
	var val int64
	//Присваиваем переменной значение2
	val = 2
	//Вызываем функцию по установке i-того бита числа в 0 или 1
	setBit(&val, 1)
	//Печатаем значение val
	fmt.Println(val)
}

//Параметры функции указатель на int64, тк мы хотим работать с участком памяти и номер бита
func setBit(number *int64, i int8) {
	*number = *number ^ (1 << i) //Устанавливаем i-тый бит в значение 0 или 1 c помощью оператора ^ (XOR)
}
