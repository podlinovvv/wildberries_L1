package main

import "fmt"

func main() {
	arr := []int{8, 11, 3, 1, 5, 4, 3, -1, 8}
	sorted := Quicksort(arr)
	fmt.Println(sorted)
}

func Quicksort(input []int) []int {
	l := len(input)
	//если длина массива меньше 2 останавливаемся
	if l < 2 {
		return input
	}

	less := make([]int, 0)
	bigger := make([]int, 0)

	//опорная точка
	pivot := input[0]
	for _, v := range input[1:] {
		if v > pivot {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(Quicksort(less), pivot)
	input = append(input, Quicksort(bigger)...)
	return input
}
