package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("ping")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("finish")
}
