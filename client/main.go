package main

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	bid, err := getBid(ctx)
	if err != nil {
		panic(err)
	}

	err = saveBidToFile(bid)
	if err != nil {
		panic(err)
	}
}

func getBid(ctx context.Context) (float64, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return 0.0, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	bid, err := io.ReadAll(res.Body)
	if err != nil {
		return 0.0, err
	}

	valor, _ := strconv.ParseFloat(string(bid[:]), 64)
	result := float64(valor)
	return result, nil
}
