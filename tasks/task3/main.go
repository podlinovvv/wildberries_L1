package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var summ int64
	mass := []int64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _,v := range mass {
		wg.Add(1)
		go func(v int64) {
			atomic.AddInt64(&summ, v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Printf("сумма квадратов: %d", summ)
}
