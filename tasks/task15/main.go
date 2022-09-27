//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

package main

import (
	"fmt"
	"strings"
)

var justString string

//проблема с этой функцией заключается в том, что строка в Golang это массив байт,
//потому присваивая глобальной переменной указатель на слайс от этого массива мы вынуждаем массив оставаться в памяти целиком

//проблемная функция
func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	*justString = v[:100]
}

//решение проблемы утечки памяти с помощью copy
func someGoodFunc(justString *string) {
	v := createHugeString(1 << 10)
	newString := make([]byte, 100)
	copy(newString, v[:100])
	*justString = string(newString)
}

func main() {
	someFunc(&justString)

	fmt.Println(justString)
}

//Допишем функцию, которая будет генерировать длинную строк заданного размера
func createHugeString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("1")
	}
	return sb.String()

}
