package usecases

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ehb1975/client_server_go/server/domain"
	"github.com/ehb1975/client_server_go/server/repository"
)

type CotacaoUseCase struct {
	repo repository.CotacaoRepository
}

func NewCotacaoUseCase(repo repository.CotacaoRepository) *CotacaoUseCase {
	return &CotacaoUseCase{
		repo: repo,
	}
}

func (uc *CotacaoUseCase) Insert() error {
	cotacao, err := BuscaCotacao()
	if err != nil {
		return err
	}
	err = uc.repo.Insert(cotacao)
	if err != nil {
		return err
	}
	return nil
}

func BuscaCotacao() (*domain.Cotacao, error) {

	req, error := http.NewRequest(http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if error != nil {
		return nil, error
	}
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)

	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c domain.Cotacao
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
