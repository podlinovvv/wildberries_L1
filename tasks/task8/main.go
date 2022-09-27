//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func main() {
	number:=int64(965)

	result1:=setBit(number,4)
	result2:=clearBit(number,4)

	fmt.Println(result1)
	fmt.Println(result2)
}

//создаём маску вида 00000100000 с единицой на месте нужного бита, после чего с помощью | изменяем на 1 только этот бит
func setBit(n int64, pos uint) int64 {
	n |= int64(1 << pos)
	return n
}
//создаём маску вида 00000100000 с единицой на месте нужного бита, после чего с помощью ^ превращаем её в обратную
//вида 11111011111  с 0 на месте нужного бита, после чего с помощью & изменяем значение этого бита на 0 в числе

func clearBit(n int64, pos uint) int64 {
	mask := ^int64(1 << pos)
	n &= mask
	return n
}

