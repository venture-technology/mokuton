package service

import "github.com/venture-technology/vtx-payout-microservice-user/internal/repository"

type PayoutService struct {
	payoutrepository repository.IPayoutRepository
}

func NewPayoutRepository(payoutrepository repository.IPayoutRepository) *PayoutService {
	return &PayoutService{
		payoutrepository: payoutrepository,
	}
}
