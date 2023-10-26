package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	//Последовательность
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	//Выведем получившееся множество
	fmt.Println(createSet(seq))
}

func createSet(seq []string) []string {
	//Создадим map, где ключами будут значения из последовательностями, а значениями пустые структуры
	m := make(map[string]struct{})
	//Создадим множество
	set := make([]string, 0)

	//Запишем все элементы последовательности в ключи map, если элемент повторится ключ просто перезапишется
	for _, val := range seq {
		m[val] = struct{}{}
	}

	//Пройдемся в цикле по ключам map и добавим их в наше множество
	for key := range m {
		set = append(set, key)
	}

	//вернем множество
	return set
}
