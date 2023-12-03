package main

import (
	"github.com/ehb1975/client_server_go/server/handlers"
	"github.com/ehb1975/client_server_go/server/infrastructure"
	"github.com/ehb1975/client_server_go/server/repository"
	"github.com/ehb1975/client_server_go/server/usecases"
	"github.com/gofiber/fiber"
)

// {"USDBRL":{"code":"USD","codein":"BRL","name":"Dólar Americano/Real Brasileiro","high":"4.937","low":"4.8684","varBid":"-0.0401","pctChange":"-0.81","bid":"4.8801",
// "ask":"4.8808","timestamp":"1701467882","create_date":"2023-12-01 18:58:02"}}
type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		CodeIn     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		Varbid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
}

func main() {

	db, err := infrastructure.GetSqliteConnection()
	if err != nil {
		return
	}
	defer infrastructure.CloseSQLDB(db)

	cotacaoRepo := repository.NewCotacaoRepositorySqlite(db)
	cotacaoService := usecases.NewCotacaoUseCase(cotacaoRepo)

	app := fiber.New()

	cotacaoHandler := handlers.NewCotacaoHandler(cotacaoService)

	app.Get("/cotacao", cotacaoHandler.Insert)
	app.Listen(":8080")

}

/*
func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cotacao, error := BuscaCotacao()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	error = userRepo.Insert(db, cotacao)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
*/
