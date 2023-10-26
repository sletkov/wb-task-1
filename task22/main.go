package main

import (
	"errors"
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
func main() {
	fmt.Println(calculateBigNumbers(big.NewInt(2<<20), "+", big.NewInt(2<<20)))
	fmt.Println(calculateBigNumbers(big.NewInt(2<<20), "-", big.NewInt(2<<20)))
	fmt.Println(calculateBigNumbers(big.NewInt(2<<20), "*", big.NewInt(2<<20)))
	fmt.Println(calculateBigNumbers(big.NewInt(2<<20), "/", big.NewInt(0)))
}

//Для решения задачи нам понадобится пакет big

//Принимаем аргументы и оператор, возвращаем результат и ошибку
func calculateBigNumbers(a *big.Int, operator string, b *big.Int) (*big.Int, error) {
	//Результат
	result := big.NewInt(0)
	//В зависимости от оператора, производим нужное действие и записываем результат
	switch operator {
	case "+":
		result.Add(a, b)
	case "-":
		result.Sub(a, b)
	case "*":
		result.Mul(a, b)
	case "/":
		//Проверка деления на 0
		if b.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("Division by 0")
		}
		result.Div(a, b)
	}

	//Возвращаем результат и nil ошибку
	return result, nil
}
