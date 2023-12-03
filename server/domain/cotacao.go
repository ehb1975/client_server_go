package domain

import "time"

type Cotacao struct {
	ID        string
	Cotacao   float64
	Timestamp time.Time
}
