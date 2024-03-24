package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	consume(channel)
}

func publish(channel chan int) {
	for i :=0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func consume(channel chan int) {
	for x := range channel {
		fmt.Println("Mensagem processada:", x)
	}
}
