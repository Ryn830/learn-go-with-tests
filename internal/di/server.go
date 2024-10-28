package di

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World!")
}

func StartServer() {
	http.HandleFunc("/hello", handleHello)

	if serverError := http.ListenAndServe(":8080", nil); serverError != nil {
		fmt.Println("Error starting server")
		fmt.Printf("Error: %q", serverError)
	}
}
