package main

import "fmt"

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false
func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

func isUnique(s string) bool {
	//Создадим map, где ключами будут символы, встречающиеся в строке, а значениями - кол-во их вхождений в строку
	m := make(map[rune]int)

	//Заполним map
	for _, val := range s {
		m[val] += 1
	}

	//Пройдемся в цикле по всем символам строки
	//Если у такого символа больше 1 вхождения, вернем false
	for _, val := range s {
		if m[val] > 1 {
			return false
		}
	}

	//Вернем true, т.к цикл сверхну не вернул false(нет символа входящего в строку более 1 раза)
	return true
}
