package main

import (
	"fmt"
	"time"
)

const timeToSleep time.Duration = 2

func main() {
	now := time.Now()
	sleep(timeToSleep * time.Second)
	fmt.Println("спали", time.Now().Sub(now))

	now = time.Now()
	sleep2(timeToSleep * time.Second)
	fmt.Println("спали", time.Now().Sub(now))
}

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func sleep2(duration time.Duration) {
	<-time.NewTimer(duration).C
}
