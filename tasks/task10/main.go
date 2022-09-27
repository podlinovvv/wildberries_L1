package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float64)

	//получаем целую часть числа и используем как ключ для map
	for _, v := range data {
		fullPart := int(v/10) * 10
		result[fullPart] = append(result[fullPart], v)
	}
	//вывод на печать всех групп
	for key, value := range result {
		fmt.Println(key, value)
	}
}
