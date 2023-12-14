package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	var req *http.Request
	var err error

	select {
	case <-ctx.Done():
		log.Println("timeout máximo de 300ms para receber o resultado do server")

	case <-time.After(300 * time.Millisecond):
		log.Println("recebido com sucesso dados do server")
		req, err = http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro na resposta do servidor: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	valor := result["dolar"]
	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", valor)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
