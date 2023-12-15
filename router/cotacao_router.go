package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/laurati/client-server-api/handler"
)

func InitializeRouter(cotacao *handler.CotacaoHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("up and running"))
	}).Methods("GET")

	router.HandleFunc("/cotacao", cotacao.GetCotacao).Methods("GET")

	return router
}
