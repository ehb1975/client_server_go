package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"

	"net/http"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
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

	http.HandleFunc("/cotacao", BuscaCotacaoHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/cotacao" {
		log.Println("not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx := r.Context()

	cotacao, error := BuscaCotacao(ctx)
	if error != nil {
		log.Println("Erro - BuscaCotacao")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	error = Insert(ctx, cotacao)
	if error != nil {
		log.Println("Erro - Insert")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao.USDBRL.Bid)

	select {
	case <-time.After(300 * time.Millisecond):
		log.Println("Request executado com sucesso")

	case <-ctx.Done():
		log.Println("Request cancelado pelo cliente")
	}
}

func Insert(ctx context.Context, cotacao *Cotacao) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Println("Erro open")
		return err
	}
	defer db.Close()

	sql := "CREATE TABLE IF NOT EXISTS cotacao (id TEXT PRIMARY KEY, cotacao REAL)"
	_, err = db.ExecContext(ctx, sql)
	if err != nil {
		log.Printf("Erro create table - %s", err)
		return err
	}

	sql = "insert into cotacao(id, cotacao) values (?,?)"
	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		log.Printf("Erro insert - %s ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid.New().String(), cotacao.USDBRL.Bid)
	if err != nil {
		log.Println("Erro exec")
		return err
	}

	return nil
}

func BuscaCotacao(ctx context.Context) (*Cotacao, error) {

	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro request")
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro do: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	body, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	var c Cotacao
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
