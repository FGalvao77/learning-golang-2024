package main

import (
	"fmt"
	"time"
)

func cont(count int) {
	for i := 0; i < count; i++ {
        	fmt.Println(i)
        	time.Sleep(1 * time.Second)
	}
	
}

func main() {
	go cont(10)
	go cont(10)
	cont(10)
}

/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 09-paralel.go
	
	
	Você terá como resposta:

		0
		0
		0
		1
		1
		1
		2
		2
		2
		3
		3
		3
		4
		4
		4
		5
		5
		5
		6
		6
		6
		7
		7
		7
		8
		8
		8
		9
		9
		9

*/
