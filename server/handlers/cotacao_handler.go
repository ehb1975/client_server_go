package handlers

import (
	"context"

	"github.com/ehb1975/client_server_go/server/domain"
	"github.com/ehb1975/client_server_go/server/usecases"
)

type CotacaoHandler struct {
	usecase *usecases.CotacaoUseCase
}

func NewCotacaoHandler(usecase *usecases.CotacaoUseCase) *CotacaoHandler {
	return &CotacaoHandler{
		usecase: usecase,
	}
}

func (h *CotacaoHandler) Insert(context *context.Context, cotacao domain.Cotacao) error {
	err := h.usecase.Insert(cotacao)
	if err != nil {
		return err
	}
	return nil
}
