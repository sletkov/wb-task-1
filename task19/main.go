package main

import (
	"fmt"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
func main() {
	fmt.Println(reverseString("главрыба"))
	fmt.Println(reverseString("привет😀"))
}

//Для работы с символами Unicode нам понадобится тип rune
func reverseString(s string) string {
	//создаем слайс рун из строки
	runeS := []rune(s)
	//Создаем переменную, где будем хранить перевернутую строку
	reverced := make([]rune, 0, len(runeS))

	//В цикле в обратном порядке проходимся по всему слайсу рун и добавляем значения в слайс reversed
	for i := len(runeS) - 1; i >= 0; i-- {
		reverced = append(reverced, runeS[i])
	}

	//Возращаем приведенный к строке слайс рун reverced
	return string(reverced)
}
