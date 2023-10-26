package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//Есть структура Dog с методом Woof
type Dog struct{}

func (d *Dog) Woof() {
	fmt.Println("Woof!")
}

//Есть структура Cat с методом Meow
type Cat struct{}

func (c *Cat) Meow(isCall bool) {
	if isCall {
		fmt.Println("Meow!")
	}
}

//Задача -свести две разные структуры к одному интерфейсу, чтобы вызывать метод с общей сигнатурой Reaction()

//Целевой интерфейс - adapter
//Сводим две разные структуры к одному интерфейсу
type AnimalAdapter interface {
	//Общая сигнатура Reaction()
	Reaction()
}

//Адаптер для собаки
type DogAdapter struct {
	*Dog //Встраиваем структуру Dog
}

//Сигнатура метода такая же как в интерфейсе AnimalAdapter. Теперь структура DogAdapter имплементирует интерфейс AnimalAdapter
func (a *DogAdapter) Reaction() {
	a.Woof() //Внутри метода вызываем метод встроенной структуры
}

//Конструктор для DogAdapter
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

//Адаптер для кошки
type CatAdapter struct {
	*Cat //Встраиваем структуру Cat
}

//Сигнатура метода такая же как в интерфейсе AnimalAdapter. Теперь структура СatAdapter имплементирует интерфейс AnimalAdapter
func (a *CatAdapter) Reaction() {
	a.Meow(true) //Внутри метода вызываем метод встроенной структуры
}

//Конструктор для CatAdapter
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	//Cоздаем CatAdapter
	cat := NewCatAdapter(&Cat{})
	//Вызываем метод Reaction
	cat.Reaction()

	//Cоздаем DogAdapter
	dog := NewDogAdapter(&Dog{})
	//Вызываем метод Reaction
	dog.Reaction()

	//У обоих структур есть метод Reaction, т.к. они имплементируют один интерфейс
}
