package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Payout struct {
	ID     uuid.UUID       `json:"id"`
	Payout json.RawMessage `json:"payout"`
	Route  string          `json:"route"`
}
