package service

import (
	"context"

	"github.com/eugenshima/Balance/internal/model"
)

type BalanceService struct {
	rps BalanceRepository
}

func NewBalanceService(rps BalanceRepository) *BalanceService {
	return &BalanceService{rps: rps}
}

type BalanceRepository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
}

func (s *BalanceService) GetAllBalances(ctx context.Context) ([]*model.User, error) {
	return s.rps.GetAll(ctx)
}
