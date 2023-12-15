package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/laurati/client-server-api/entity"
)

type CotacaoRepo struct {
	db *sql.DB
}

func NewCotacaoRepo(db *sql.DB) *CotacaoRepo {
	return &CotacaoRepo{
		db: db,
	}
}

func (c *CotacaoRepo) CreateCotacao(ctx context.Context, cotacao *entity.Cotacao) error {

	stmt, err := c.db.PrepareContext(ctx, "INSERT INTO cotacoes (valor) VALUES (?)")
	if err != nil {
		log.Printf("Erro ao preparar a declaração SQL: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, cotacao.Bid)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalln("o timeout máximo para persistir os dados no banco deverá ser de 10ms")
		}
		log.Printf("Erro ao inserir cotação no banco de dados: %v", err)
		return err
	}
	log.Println("cotacao inserida no db com sucesso")

	return nil
}
