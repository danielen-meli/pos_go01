package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// func que demostra com uma msg que o servidor subiu na rota "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go no AR!")
	})

	// cria o endpoint 'cotação' que ainda não tras a resposta, por isso um 'mock' de como seria a resposta
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		apiToConsume := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

		// Fazer a requisição com um timeout configurado no contexto
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		// ja usando o ctx, configura como será a achamada pra api de cotacao
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiToConsume, nil)
		if err != nil {
			http.Error(w, "Erro ao acessar a API de cotação", http.StatusInternalServerError)
			fmt.Printf("Error ao fazer a requisição da cotação: %v/n", err)
			return
		}

		client := http.Client{}     // instancia client
		resp, err := client.Do(req) //faz o get na api pra de fato consumir o endpoint
		if err != nil {
			http.Error(w, "Erro ao requisitar API de cotação", http.StatusInternalServerError)
			fmt.Printf("Erro ao fazer requisição para API: %v\n", err)
			return
		}
		defer resp.Body.Close() // tem que ter o defer close, e pode ser aqui antes de ler o body pra não esquecer depois

		body, err := io.ReadAll(resp.Body)

		// onde será armazenado a resp em json
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			http.Error(w, "Erro ao decodificar JSON da API", http.StatusInternalServerError)
			fmt.Printf("Erro ao fazer unmarshal do JSON: %v\n", err)
			return
		}

		bid := result["USDBRL"].(map[string]interface{})["bid"].(string)

		jsonResponse := fmt.Sprintf(`{"bid": "%s"}`, bid)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(jsonResponse))
		if err != nil {
			fmt.Printf("Erro na resposta do JSON %v\n", err)
		}
	})

	fmt.Println("Servidor na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error : %v\n", err)
	}
}
