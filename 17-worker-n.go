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

  numWorkers := 100

  for i := 0; i < numWorkers; i++ {
    go worker(i, channel)
  }

  for i := 0; i < 1000; i++ {
    channel <- i
  }
}