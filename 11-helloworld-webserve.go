package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nHello, World!\n\n"))
}


/*

	Abra um terminal ou o shell e digite o seguinte comando:
	
		go run 11-helloworld-webserve.go

	Abra um novo terminal e digite o seguinte comando:
	
		curl localhost:3000
	
	
	Você terá como resposta:
	
	 Hello, World!

*/
