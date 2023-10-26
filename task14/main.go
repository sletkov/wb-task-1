package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.
func main() {
	fmt.Println(getType(1))
	fmt.Println(getType("string"))
	fmt.Println(getType(true))
	fmt.Println(getType(make(chan int)))
}

//Для определения типа воспользуемся пакетом reflect, а конкретно функцией TypeOf
func getType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
