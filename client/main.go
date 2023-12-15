package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("o client terá um timeout máximo de 300ms para receber o resultado do server")
		}
		log.Fatal(err)
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

	// file 1
	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", valor)), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// file 2
	file, err := os.Create("fileCotacoes.txt")
	if err != nil {
		log.Printf("erro %v ao criar arquivo %v\n", err.Error(), file)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", valor))
	if err != nil {
		log.Printf("erro %v ao escrever no arquivo %v\n", err.Error(), file)
	}

	fmt.Printf("Dólar: %s\n", valor)
}
