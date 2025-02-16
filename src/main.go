package main

import (
	"fmt"
	"net/http"
)

func main() {
	// func que será chamada ao acessar o endpoint "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go")
	})

	fmt.Println("Servidor na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error : %v\n", err)
	}
}
