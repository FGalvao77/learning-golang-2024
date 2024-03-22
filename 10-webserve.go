package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", nil)
}

/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 10-webserve.go

	Abra um novo terminal e digite o seguinte comando:
	
		curl localhost:3000
	
	
	Você terá como resposta:
	
	 404 page not found
	 
	 Ou seja, página não encontrada.

*/
