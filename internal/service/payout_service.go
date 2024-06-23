package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/vtx-payout-microservice-user/internal/repository"
	"github.com/venture-technology/vtx-payout-microservice-user/models"
)

type PayoutService struct {
	payoutrepository repository.IPayoutRepository
}

func NewPayoutService(payoutrepository repository.IPayoutRepository) *PayoutService {
	return &PayoutService{
		payoutrepository: payoutrepository,
	}
}

func (s *PayoutService) CreatePayout(ctx context.Context, payout *models.Payout) error {

	id, err := uuid.NewV7()

	if err != nil {
		return err
	}

	payout.ID = id

	return s.payoutrepository.CreatePayout(ctx, payout)

}

func (s *PayoutService) ReadPayout(ctx context.Context, id *string) (*models.Payout, error) {

	uuid, err := uuid.Parse(*id)

	if err != nil {
		return nil, err
	}

	return s.payoutrepository.ReadPayout(ctx, &uuid)
}

func (s *PayoutService) DeletePayout(ctx context.Context, id *string) error {

	uuid, err := uuid.Parse(*id)

	if err != nil {
		return err
	}

	return s.payoutrepository.DeletePayout(ctx, &uuid)
}
