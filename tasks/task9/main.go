package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//канал между первой и второй горутинами
	ch1 := make(chan int)
	//канал между второй и третьей горутинами
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	//горутина итерирует по слайсу и пишет числа в канал ch1
	go func() {
		defer wg.Done()
		defer close(ch1)

		for i := range arr {
			ch1 <- arr[i]
		}

	}()

	//горутина читает числа из канала ch1, умножает их на 2 и отправляет дальше в канал ch2
	go func() {
		defer wg.Done()
		defer close(ch2)

		//когда канал закроется цикл прервётся
		for q := range ch1 {
			q *= 2
			ch2 <- q
		}
	}()

	//горутина читет числа из канала ch2 и выводит их на печать
	go func() {
		defer wg.Done()
		//когда канал закроется цикл прервётся
		for result := range ch2 {
			fmt.Println(result)
		}
	}()

	//ожидаем завершения всех горутин
	wg.Wait()
}
