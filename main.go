package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/laurati/client-server-api/handler"
	"github.com/laurati/client-server-api/repository"
	"github.com/laurati/client-server-api/router"
	"github.com/laurati/client-server-api/server"
)

func main() {

	log.Printf("Server Port: 8080")
	log.Println("Starting process API...")

	db, err := sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		log.Printf("error opening database: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewCotacaoRepo(db)
	handler := handler.NewCotacaoHandler(repo)
	router := router.InitializeRouter(handler)

	server := server.NewServer(":8080", router)
	server.Start()

}
