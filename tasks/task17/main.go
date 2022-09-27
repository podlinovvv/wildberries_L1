package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 11, 14, 17, 18, 26, 34, 45, 63, 134}
	binarySearch(arr, 14)
	binarySearch(arr, 13)

}

func binarySearch(slice []int, number int) {
	left := 0
	right := len(slice) - 1

	for left <= right {
		median := (left + right) / 2

		if slice[median] == number {
			fmt.Printf("число %d найдено по индексу: %d\n", number, median)
			return
		}
		if slice[median] > number {
			right = median - 1
			continue
		}
		left = median + 1

	}
	fmt.Printf("число %d не найдено в массиве\n", number)
	return
}
