package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	//запускаем в цикле горутины для каждого числа, отслеживаем их активность через waitgroup
	for _,v := range arr {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			wg.Done()
		}(v)
	}
	//ждём окончания всех горутин
	wg.Wait()
}
