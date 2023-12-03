package repo

import (
	"database/sql"

	"github.com/ehb1975/client_server_go/domain"
	_ "github.com/mattn/go-sqlite3"
)

type CotacaoRepositorySqlite struct {
	db *sql.DB
}

func NewCotacaoRepositorySqlite(db *sql.DB) *CotacaoRepositorySqlite {
	return &CotacaoRepositorySqlite{
		db: db,
	}
}

func (r *CotacaoRepositorySqlite) Insertcotacao(cotacao domain.Cotacao) error {
	stmt, err := db.Prepare("insert into cotacao(id, cotacao, timestamp) values ($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cotacao.ID, cotacao.Cotacao, cotacao.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
