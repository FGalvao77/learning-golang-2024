package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) Walk() {
	fmt.Println(p.name + " is an Analyst & Data Scientist.")
	fmt.Println("He is", p.age, "years old!")
}

func main() {
	var person Person
	person.name = "Fernando"
	person.age = 47

	person.Walk()
}

/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 07-datatype2.go
	
	
	Você terá como resposta:

		Fernando is an Analyst & Data Scientist.
		He is 47 years old! 

*/
