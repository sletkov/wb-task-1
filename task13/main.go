package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	//В go можно менять местами значения переменной без создания третьей переменной
	a, b := 1, 2
	a, b = b, a

	fmt.Println(a, b)
}
