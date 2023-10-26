package main

import "fmt"

// Удалить i-ый элемент из слайса.
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(deleteElementByIdx(slice1, 7))
	fmt.Println(deleteElementByIdx2(slice2, 7))
}

//Выполняется за линейное время и сохраняет порядок элементов
func deleteElementByIdx(slice []int, idx int) []int {
	//Возвращаем объединение двух слайсов
	//Первый слайс - элементы начиная с первого до того, который хотим удалить(не включая его)
	//Второй слайс - элементы начиная со следующего после того, который мы хотим удалить до конца.
	return append(slice[:idx], slice[idx+1:]...)
}

//Выполняется за константное время, но не сохраняет порядок
func deleteElementByIdx2(slice []int, idx int) []int {
	//ставим на место элемента, которой хотим удалить - последний элемент
	slice[idx] = slice[len(slice)-1]
	//возвращаем усеченный на 1 элемент слайс
	return slice[:len(slice)-1]
}
