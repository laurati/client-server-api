package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

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

	select {
	case <-ctx.Done():
		log.Println("timeout máximo para inserir dados no db é 10ms")

	case <-time.After(100 * time.Millisecond):
		_, err := c.db.ExecContext(ctx, "INSERT INTO cotacoes (valor) VALUES (?)", cotacao.Bid)
		if err != nil {
			log.Printf("Erro ao inserir cotação no banco de dados: %v", err)
			return err
		}
		log.Println("cotacao inserida no db com sucesso")
	}

	return nil
}
