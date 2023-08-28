package handlers

import (
	"context"

	"github.com/eugenshima/Balance/internal/model"
)

type BalanceHandler struct {
	srv BalanceService
}

func NewBalancehandler(srv BalanceService) *BalanceHandler {
	return &BalanceHandler{srv: srv}
}

type BalanceService interface {
	GetAllBalances(ctx context.Context) ([]*model.User, error)
}

func (h *BalanceHandler) GetAllBalances(ctx context.Context)
