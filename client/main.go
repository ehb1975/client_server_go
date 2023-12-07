package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	bid, err := getBid(ctx)
	if err != nil {
		log.Printf("Erro getBid")
		panic(err)
	}

	err = saveBidToFile(bid)
	if err != nil {
		panic(err)
	}
}

func saveBidToFile(bid string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}

	_, err = file.WriteString("Dollar: " + bid)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func getBid(ctx context.Context) (string, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("Erro NewRequest")
		return "", err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println("Erro do")
		return "", err
	}
	defer res.Body.Close()

	bid, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ReadAll")
		return "", err
	}

	return string(bid[:]), nil
}
