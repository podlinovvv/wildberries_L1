package main

import (
	"fmt"
	"reflect"
)

func main() {
	//int
	reflectChecker(int8(1))
	reflectChecker(uint64(99999))
	//string
	reflectChecker("строка")
	//bool
	reflectChecker(true)
	reflectChecker(false)
	//channel
	channel := make(chan int)
	reflectChecker(channel)

	//int
	sprintChecker(int8(1))
	sprintChecker(uint64(99999))
	//string
	sprintChecker("строка")
	//bool
	sprintChecker(true)
	sprintChecker(false)
	//channel
	channel = make(chan int)
	sprintChecker(channel)
}

func reflectChecker(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}

func sprintChecker(value interface{}) {
	fmt.Println(fmt.Sprintf("%T", value))
}
