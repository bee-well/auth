package domain

import "time"

type Token struct {
	ID     string
	Issued time.Time
}
