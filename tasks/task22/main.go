package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создание длинных чисел
	number1, number2 := new(big.Float), new(big.Float)

	number1.SetString("999999999999999999999999999999999999999999999999999999999999999.9999999999")
	number2.SetString("777777777777777777777777777777777777777777777777777777777777777.7777777777")

	//форматирование чисел
	n1 := number1.Text('f', 5)
	n2 := number2.Text('f', 5)

	//проведение операций на числах
	add := addition(number1, number2).Text('f', 5)
	sub := subtraction(number1, number2).Text('f', 5)
	multi := multiplication(number1, number2).Text('f', 5)
	div := division(number1, number2).Text('f', 5)

	//вывод результатов
	fmt.Printf("\nfirst number = %s\n\n", n1)
	fmt.Printf("second number  = %s\n\n", n2)
	fmt.Printf("addition = %s\n\n", add)
	fmt.Printf("subtraction = %s\n\n", sub)
	fmt.Printf("multiplication = %s\n\n", multi)
	fmt.Printf("division = %s", div)
}

//сложение
func addition(number1 *big.Float, number2 *big.Float) *big.Float {
	res := new(big.Float)
	return res.Add(number1, number2)
}

//вычитание
func subtraction(number1 *big.Float, number2 *big.Float) *big.Float {
	res := new(big.Float)
	return res.Sub(number1, number2)
}

//умножение
func multiplication(number1 *big.Float, number2 *big.Float) *big.Float {
	res := new(big.Float)
	return res.Mul(number1, number2)
}

//деление
func division(number1 *big.Float, number2 *big.Float) *big.Float {
	res := new(big.Float)
	return res.Quo(number1, number2)
}
