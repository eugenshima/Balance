// Package handlers contains gRPC methods
package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"
	proto "github.com/eugenshima/Balance/proto"
	"github.com/sirupsen/logrus"

	vld "github.com/go-playground/validator"
	"github.com/google/uuid"
)

// BalanceHandler struct represents a balance handler
type BalanceHandler struct {
	srv BalanceService
	vl  *vld.Validate
	proto.UnimplementedUserServiceServer
}

// NewBalancehandler Creates a new BalanceHandler
func NewBalancehandler(srv BalanceService, vl *vld.Validate) *BalanceHandler {
	return &BalanceHandler{
		srv: srv,
		vl:  vl,
	}
}

// BalanceService interface represents service methods
type BalanceService interface {
	GetAllBalances(ctx context.Context) ([]*model.User, error)
	UpdateBalance(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateBalance(ctx context.Context, user *model.User) error
	DeleteBalance(ctx context.Context, userID string) error
}

// Validate func validates your model
func (h *BalanceHandler) Validate(i interface{}) error {
	if err := h.vl.Struct(i); err != nil {
		logrus.Errorf("Struct: %v", err)
		return fmt.Errorf("struct: %w", err)
	}
	return nil
}

// UpdateUserBalance function updates the user balance information
func (h *BalanceHandler) UpdateUserBalance(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	ID, err := uuid.Parse(req.User.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.User.ID": req.User.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	user := &model.User{
		ID:       ID,
		Username: req.User.Username,
		Balance:  float64(req.User.Balance),
	}
	err = h.Validate(user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user}).Errorf("Validate: %v", err)
		return nil, fmt.Errorf("Validate: %w", err)
	}
	err = h.srv.UpdateBalance(ctx, user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user}).Errorf("UpdateBalance: %v", err)
		return nil, fmt.Errorf("UpdateBalance: %w", err)
	}
	return &proto.UserUpdateResponse{}, nil
}

// GetUserByID function returns a user with the given ID
func (h *BalanceHandler) GetUserByID(ctx context.Context, req *proto.UserGetByIDRequest) (*proto.UserGetByIDResponse, error) {
	result, err := h.srv.GetUserByID(ctx, req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.ID": req.ID}).Errorf("GetUserByID: %v", err)
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	err = h.Validate(result)
	if err != nil {
		logrus.WithFields(logrus.Fields{"result": result}).Errorf("Validate: %v", err)
		return nil, fmt.Errorf("Validate: %w", err)
	}
	user := &proto.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Balance:  float64(result.Balance),
	}
	return &proto.UserGetByIDResponse{User: user}, nil
}

// CreateUserBalance function creates a new user balance
func (h *BalanceHandler) CreateUserBalance(ctx context.Context, req *proto.CreateBalanceRequest) (*proto.CreateBalanceResponse, error) {
	user := &model.User{
		ID:       uuid.New(),
		Username: req.User.Username,
		Balance:  float64(req.User.Balance),
	}
	err := h.Validate(user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user}).Errorf("Validate: %v", err)
		return nil, fmt.Errorf("Validate error: %w", err)
	}
	err = h.srv.CreateBalance(ctx, user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user}).Errorf("CreateBalance: %v", err)
		return nil, fmt.Errorf("CreateBalance: %w", err)
	}
	return &proto.CreateBalanceResponse{}, nil
}
