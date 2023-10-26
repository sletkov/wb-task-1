package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

//Объявляем структуру Point
type Point struct {
	x, y int
}

//Конструктор структуры Point
func NewPoint(x, y int) Point {
	return Point{
		x,
		y,
	}
}

func main() {
	//Создаем две точки с разными координатами
	p1 := NewPoint(1, 1)
	p2 := NewPoint(0, 1)

	//Выводим расстояние между двумя точками, которое возвращает метод getDistance
	fmt.Println(getDistance(p1, p2))
}

func getDistance(p1, p2 Point) float64 {
	//Воспользуемся формулой нахождения расстояния между двумя точками d = sqrt( (x1 - x2)^2 + (y1 - y2)^2 )
	return math.Sqrt(math.Pow((float64(p1.x-p2.x)), 2) + math.Pow(float64(p1.y-p2.y), 2))
}
