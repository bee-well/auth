package domain

import "time"

type Token struct {
	ID     int64
	Issued time.Time
}
