package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

//количество воркеров
const N int = 5

func main() {
	//seed для генерации случайных данных
	rand.Seed(time.Now().UnixNano())

	channel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(N)

	//контекст уведомит Writer-а о завершении работы программы
	notifyContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for i := 0; i < N; i++ {
		go Worker(&wg, channel)
	}
	Writer(notifyContext, channel)
	wg.Wait()

}

func Writer(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("ctrl+c pressed")
			return
		default:
			data := rand.Intn(1000)
			ch <- data
		}
	}
}

func Worker(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for data := range c {
		fmt.Println(data)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("worker finished")
}
