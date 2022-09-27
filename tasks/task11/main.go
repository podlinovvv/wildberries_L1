package main

import "fmt"

func main() {
	arr1 := []int{13, 9, 5, 2, 7, 14, 9, 2, 5, 4, 11, 3, 6, 2, 8, 2, 8}
	arr2 := []int{14, 10, 1, 8, 6, 13, 2, 5, 3, 8, 14, 12, 8, 2, 6, 2, 7, 8, 9}

	fmt.Println(intersect(&arr1, &arr2))

}

func intersect(arr1, arr2 *[]int) []int {
	var result []int
	set := make(map[int]struct{})

	//из одного массива записываем в map
	for _, v := range *arr1 {
		set[v] = struct{}{}
	}
	//проходим по второму и проверяем есть ли эти числа в первом. Если есть, пишем в результирующий массив
	for _, v := range *arr2 {
		_, ok := set[v]
		if ok {
			result = append(result, v)
		}
	}
	return result
}
