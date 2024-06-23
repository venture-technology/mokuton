package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/vtx-payout-microservice-user/models"
)

type IPayoutRepository interface {
	CreatePayout(ctx context.Context, payout *models.Payout) error
	ReadPayout(ctx context.Context, id *uuid.UUID) (*models.Payout, error)
	DeletePayout(ctx context.Context, id *uuid.UUID) error
}

type PayoutRepository struct {
	db *sql.DB
}

func NewPayoutRepository(conn *sql.DB) *PayoutRepository {
	return &PayoutRepository{
		db: conn,
	}
}

func (r *PayoutRepository) CreatePayout(ctx context.Context, payout *models.Payout) error {

	sqlQuery := `INSERT INTO payouts (id, payout, route) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(sqlQuery, payout.ID, payout.Payout, payout.Route)

	return err

}

func (r *PayoutRepository) ReadPayout(ctx context.Context, id *uuid.UUID) (*models.Payout, error) {

	sqlQuery := `SELECT id, payout, route FROM payouts WHERE id = $1 LIMIT 1`

	var payout models.Payout

	err := r.db.QueryRow(sqlQuery, *id).Scan(
		&payout.ID,
		&payout.Payout,
		&payout.Route,
	)

	if err != nil {
		return nil, err
	}

	return &payout, nil

}

func (r *PayoutRepository) DeletePayout(ctx context.Context, id *uuid.UUID) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM payouts WHERE id = $1", *id)
	return err

}
