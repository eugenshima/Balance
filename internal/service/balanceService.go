// Package service provides a set of functions, which include business-logic in it
package service

import (
	"context"

	"github.com/eugenshima/Balance/internal/model"

	"github.com/google/uuid"
)

// BalanceService struct represents a Balance Service
type BalanceService struct {
	rps BalanceRepository
}

// NewBalanceService function creates a new Balance Service
func NewBalanceService(rps BalanceRepository) *BalanceService {
	return &BalanceService{rps: rps}
}

// BalanceRepository represents a Balance Repository methods
type BalanceRepository interface {
	GetAll(ctx context.Context) ([]*model.Balance, error)
	UpdateBalance(ctx context.Context, user *model.Balance) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*model.Balance, error)
	CreateBalance(ctx context.Context, user *model.Balance) error
	DeleteBalance(ctx context.Context, userID uuid.UUID) error
}

// GetAllBalances function returns Get All repository method
func (s *BalanceService) GetAllBalances(ctx context.Context) ([]*model.Balance, error) {
	return s.rps.GetAll(ctx)
}

// UpdateBalance function returns Update repository method
func (s *BalanceService) UpdateBalance(ctx context.Context, user *model.Balance) error {
	return s.rps.UpdateBalance(ctx, user)
}

// GetUserByID function returns Get By ID repository method
func (s *BalanceService) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.Balance, error) {
	return s.rps.GetUserByID(ctx, userID)
}

// CreateBalance function returns Create repository method
func (s *BalanceService) CreateBalance(ctx context.Context, user *model.Balance) error {
	return s.rps.CreateBalance(ctx, user)
}

// DeleteBalance function returns Delete repository method
func (s *BalanceService) DeleteBalance(ctx context.Context, userID uuid.UUID) error {
	return s.rps.DeleteBalance(ctx, userID)
}
