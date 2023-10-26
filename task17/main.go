package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(binarySearch(slice, 10))
	fmt.Println(binarySearch(slice, 55))
}

//Бинарный поиск выполняется за время O(LogN)
//Необходимо, чтобы слайс был отсортирован
func binarySearch(slice []int, val int) int {
	//Ставим левый указатель на первый индекс слайса
	leftPointer := 0

	//Ставим правый указатель на последний индекс слайса
	rightPointer := len(slice) - 1

	//Пока указатели не встретились, будем их двигать навстречу друг другу
	for leftPointer <= rightPointer {
		//Найдем серединное значение и серединный указатель
		midPointer := int((leftPointer + rightPointer) / 2)
		midValue := slice[midPointer]

		//Если серединное значение равно искомому, вернем индекс этого значения
		if midValue == val {
			return midPointer
		} else if midValue < val {
			//Если серединное значение меньше искомого, сдвигаем левый указатель на место
			//индекса, следующего после индекса серединного значения
			leftPointer = midPointer + 1
		} else {
			//Если серединное значение больше искомого, сдвигаем правый указатель на место
			//индекса, следующего до индекса серединного значения
			rightPointer = midPointer - 1
		}
	}

	//Вернем -1, если не нашли элемент
	return -1
}
