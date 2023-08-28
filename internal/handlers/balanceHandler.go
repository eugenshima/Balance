package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"
	proto "github.com/eugenshima/Balance/proto"
	"github.com/google/uuid"
)

type BalanceHandler struct {
	srv BalanceService
	proto.UnimplementedUserServiceServer
}

func NewBalancehandler(srv BalanceService) *BalanceHandler {
	return &BalanceHandler{srv: srv}
}

type BalanceService interface {
	GetAllBalances(ctx context.Context) ([]*model.User, error)
	UpdateBalance(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateBalance(ctx context.Context, user *model.User) error
}

// UpdateUserBalance function updates the user balance information
func (h *BalanceHandler) UpdateUserBalance(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	ID, err := uuid.Parse(req.User.ID)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	user := &model.User{
		ID:       ID,
		Username: req.User.Username,
		Balance:  float64(req.User.Balance),
	}
	err = h.srv.UpdateBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("UpdateBalance: %w", err)
	}
	return &proto.UserUpdateResponse{}, nil
}

// GetUserByID function returns a user with the given ID
func (h *BalanceHandler) GetUserByID(ctx context.Context, req *proto.UserGetByIDRequest) (*proto.UserGetByIDResponse, error) {
	result, err := h.srv.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	user := &proto.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Balance:  float64(result.Balance),
	}
	return &proto.UserGetByIDResponse{User: user}, nil
}

func (h *BalanceHandler) CreateUserBalance(ctx context.Context, req *proto.CreateBalanceRequest) (*proto.CreateBalanceResponse, error) {
	user := &model.User{
		ID:       uuid.New(),
		Username: req.User.Username,
		Balance:  float64(req.User.Balance),
	}
	err := h.srv.CreateBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("CreateBalance: %w", err)
	}
	return &proto.CreateBalanceResponse{}, nil
}
