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
	UpdateBalance(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateBalance(ctx context.Context, user *model.User) error
}

func (s *BalanceService) GetAllBalances(ctx context.Context) ([]*model.User, error) {
	return s.rps.GetAll(ctx)
}

func (s *BalanceService) UpdateBalance(ctx context.Context, user *model.User) error {
	return s.rps.UpdateBalance(ctx, user)
}

func (s *BalanceService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	return s.rps.GetUserByID(ctx, userID)
}

func (s *BalanceService) CreateBalance(ctx context.Context, user *model.User) error {
	return s.rps.CreateBalance(ctx, user)
}
