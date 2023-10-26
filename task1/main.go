package main

import "fmt"

//Структура Human с полями name и age
type Human struct {
	name string
	age  int
}

//Методы структуры Human
func (h Human) PrintName() {
	fmt.Println("Human name is ", h.name)
}

func (h Human) PrintAge() {
	fmt.Println("Human age is ", h.age)
}

func (h Human) SayHello() {
	fmt.Println("Hello")
}

//Встраиваем структуру Human в структуру Action
type Action struct {
	Human
}

func main() {
	a := &Action{}
	//Теперь мы можем использовать методы Human у структуры Action
	a.SayHello()
	a.PrintAge()
}
