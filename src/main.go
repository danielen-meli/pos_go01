package main

import (
	"fmt"
	"net/http"
)

func main() {
	// func que será chamada ao acessar o endpoint "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go no AR!")
	})

	// cria o endpoint 'cotação' que ainda não tras a resposta, por isso um 'mock' de como seria a resposta
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		// mock da resposta fixo só pra ver qdo chama /cotacao
		jsonResponse := `{"bid": "6.00"}`

		// como é json tem q ser o Header de content type
		w.Header().Set("Content-Type", "application/json")

		// json na resposta já mostrando se tiver erro
		_, err := w.Write([]byte(jsonResponse))
		if err != nil {
			fmt.Printf("Error na resposta do Json %v\n", err)
		}
	})

	fmt.Println("Servidor na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error : %v\n", err)
	}
}
