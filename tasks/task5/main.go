package main

import (
	"fmt"
	"sync"
	"time"
)

//задаём время работы программы N в секундах
const Duration time.Duration = 3

func main() {
	//channel creation
	ch := make(chan int)
	var wg sync.WaitGroup

	//горутина, которая читает из канала
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("reading goroutine end")
				return
			}
		}
	}()

	//горутина, которая пишет в канал
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			ch <- 1
			select {
			//завершение цикла по истечении времени
			case <-time.After(Duration * time.Second):
				close(ch)
				fmt.Println("writing goroutine end")
				return
			}
		}
	}()
	//Ожидание окончания работы всех горутин
	wg.Wait()
}
