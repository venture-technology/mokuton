package controller

import "github.com/venture-technology/vtx-payout-microservice-user/internal/service"

type PayoutController struct {
	payoutservice *service.PayoutService
}

func NewPayoutController(payoutservice *service.PayoutService) *PayoutController {
	return &PayoutController{
		payoutservice: payoutservice,
	}
}
