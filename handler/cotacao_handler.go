package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/laurati/client-server-api/entity"
	"github.com/laurati/client-server-api/repository"
	_ "github.com/mattn/go-sqlite3"
)

type CotacaoHandler struct {
	Repo *repository.CotacaoRepo
}

func NewCotacaoHandler(Repo *repository.CotacaoRepo) *CotacaoHandler {
	return &CotacaoHandler{
		Repo: Repo,
	}
}

func (h *CotacaoHandler) GetCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := getCotacao(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel = context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer cancel()
	err = h.Repo.CreateCotacao(ctx, cotacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"dolar": cotacao.Bid}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Erro ao serializar resposta JSON: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getCotacao(ctx context.Context) (*entity.Cotacao, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms")
			return nil, err
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRateMap map[string]entity.Cotacao
	err = json.Unmarshal(body, &exchangeRateMap)
	if err != nil {
		log.Println("error unmarshal:", err)
		return nil, err
	}

	var key string
	for chave := range exchangeRateMap {
		key = chave
		break
	}

	exchangeRate := exchangeRateMap[key]

	cotacao := entity.Cotacao{
		Bid: exchangeRate.Bid,
	}

	return &cotacao, nil
}
