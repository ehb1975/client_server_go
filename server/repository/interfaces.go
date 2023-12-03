package repository

import "github.com/ehb1975/client_server_go/server/domain"

type CotacaoRepository interface {
	Insert(cotacao *domain.Cotacao) error
}
