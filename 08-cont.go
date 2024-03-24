package main

import (
	"fmt"
	"time"
)

func cont(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
        	fmt.Println(i)
	}
	
}

func main() {
	cont(10)
}

/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 08-cont.go
	
	
	Você terá como resposta:

		0
		1
		2
		3
		4
		5
		6
		7
		8
		9

*/
