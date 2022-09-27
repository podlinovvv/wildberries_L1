package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//используется WaitGroup для синхронизации выполнения горутин
	//и atomic для потокобезопасного увеличения значения счётчика
	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}