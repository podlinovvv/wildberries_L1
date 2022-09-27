package main

import (
	"errors"
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice, err := remove(slice, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)
}

func remove(slice []int, i int) ([]int, error) {
	if i >= len(slice) || i < 0 || slice == nil {
		return nil, errors.New("invalid index for this array")
	}
	return append(slice[:i], slice[i+1:]...), nil
}
