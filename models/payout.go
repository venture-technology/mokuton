package models

import "github.com/google/uuid"

type Payout struct {
	ID     uuid.UUID `json:"id"`
	Payout string    `json:"payout"`
	Route  string    `json:"route"`
}
