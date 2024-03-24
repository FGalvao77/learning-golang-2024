package main

import "fmt"

// T1
func main() {
	channel := make(chan string)
	
	// T2
	go func() {
		channel <- "Golang Conference! - T2"
	}()

	// T1
	msg := <- channel 
	fmt.Println(msg)
}
