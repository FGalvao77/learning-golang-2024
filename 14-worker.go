package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", workerId, x)
		}
	}

func main() {
	channel := make(chan int)
	go worker(1, channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}
}
