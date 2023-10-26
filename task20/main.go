package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}

func reverseWords(s string) string {
	//Слайс слов
	words := strings.Fields(s)
	//Слайс слов в обратном порядке
	reversedWords := make([]string, len(words))

	//В цикле в обратном порядке пройдемся по слайс слов
	//Запишем эти слова в слайс слов в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords[len(words)-i-1] = words[i]
	}

	//Вернем слайс слов в обратном порядке, превращенный в строку
	return strings.Join(reversedWords, " ")
}
