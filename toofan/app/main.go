package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :8003")
	if err := http.ListenAndServe(":8003", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
