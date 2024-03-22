package main

import "fmt"

type Person struct {
	name string
	age int
	adult bool
}

func main() {
	var citizen1 Person
	citizen1.name = "Fernando"
	citizen1.age = 47
	citizen1.adult = true
        
	var citizen2 Person
	citizen2.name = "Kátia"
	citizen2.age = 44
	citizen2.adult = true

	var citizen3 Person
	citizen3.name = "Sara"
	citizen3.age = 24
	citizen3.adult = true

	var citizen4 Person
	citizen4.name = "Eloah"
	citizen4.age = 16
	citizen4.adult = false


	println("\nPerson 1")
	fmt.Println(citizen1.name)
	fmt.Println(citizen1.age)
	fmt.Println(citizen1.adult)
	
	println("\nPerson 2")
	fmt.Println(citizen2.name)
	fmt.Println(citizen2.age)
	fmt.Println(citizen2.adult)
	
	println("\nPerson 3")
	fmt.Println(citizen3.name)
	fmt.Println(citizen3.age)
	fmt.Println(citizen3.adult)
	
	println("\nPerson 4")
	fmt.Println(citizen4.name)
	fmt.Println(citizen4.age)
	fmt.Println(citizen4.adult)
}


/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 06-datatype.go
	
	
	Você terá como resposta:
	
		Person 1
		Fernando
		47
		true
		
		Person 2
		Kátia
		44
		true
		
		Person 3
		Sara
		24
		true
		
		Person 4
		Eloah
		47
		false

*/
