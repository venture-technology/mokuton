package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/vtx-payout-microservice-user/models"
)

type IPayoutRepository interface {
	CreatePayout(ctx context.Context, payout *models.Payout) error
	ReadPayout(ctx context.Context, route *string) (*models.Payout, error)
	DeletePayout(ctx context.Context, uuid *uuid.UUID) error
}

type PayoutRepository struct {
	db *sql.DB
}

func NewPayoutRepository(conn *sql.DB) *PayoutRepository {
	return &PayoutRepository{
		db: conn,
	}
}
