package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("a:%d, b:%d\n", a, b)

	a, b = b, a
	fmt.Printf("a:%d, b:%d", a, b)
}