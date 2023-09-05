// Package handlers contains gRPC methods
package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"
	proto "github.com/eugenshima/Balance/proto"

	vld "github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// BalanceHandler struct represents a balance handler
type BalanceHandler struct {
	srv BalanceService
	vl  *vld.Validate
	proto.UnimplementedBalanceServiceServer
}

// NewBalancehandler Creates a new BalanceHandler
func NewBalancehandler(srv BalanceService, vl *vld.Validate) *BalanceHandler {
	return &BalanceHandler{
		srv: srv,
		vl:  vl,
	}
}

//go:generate /home/yauhenishymanski/work/bin/mockery --name=BalanceService --case=underscore --output=./mocks

// BalanceService interface represents service methods
type BalanceService interface {
	GetAllBalances(ctx context.Context) ([]*model.Balance, error)
	UpdateBalance(ctx context.Context, user *model.Balance) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*model.Balance, error)
	CreateBalance(ctx context.Context, user *model.Balance) error
	DeleteBalance(ctx context.Context, userID uuid.UUID) error
}

// CustomIDValidaion func validates your variables
func (h *BalanceHandler) CustomIDValidaion(ctx context.Context, i interface{}) error {
	err := h.vl.VarCtx(ctx, i, "required")
	if err != nil {
		return fmt.Errorf("VarCtx: %w", err)
	}
	err = h.vl.VarCtx(ctx, i, "uuid")
	if err != nil {
		return fmt.Errorf("VarCtx: %w", err)
	}
	return nil
}

// UpdateUserBalance function updates the user balance information
func (h *BalanceHandler) UpdateUserBalance(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	err := h.CustomIDValidaion(ctx, req.Balance.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.Balance.User_ID": req.Balance.ProfileID}).Errorf("CustomValidate: %v", err)
		return nil, fmt.Errorf("validate: %w", err)
	}
	ID, err := uuid.Parse(req.Balance.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.User.Balance_ID": req.Balance.ProfileID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	user := &model.Balance{
		BalanceID: ID,
		Balance:   req.Balance.Balance,
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
	err := h.CustomIDValidaion(ctx, req.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ProfileID": req.ProfileID}).Errorf("Validate: %v", err)
		return nil, fmt.Errorf("validate: %w", err)
	}
	ID, err := uuid.Parse(req.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.ProfileID": req.ProfileID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	result, err := h.srv.GetUserByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.Balance_ID": req.ProfileID}).Errorf("GetUserByID: %v", err)
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}

	balance := &proto.Balance{
		ProfileID: result.BalanceID.String(),
		Balance:   result.Balance,
	}
	return &proto.UserGetByIDResponse{Balance: balance}, nil
}

// CreateUserBalance function creates a new user balance
func (h *BalanceHandler) CreateUserBalance(ctx context.Context, req *proto.CreateBalanceRequest) (*proto.CreateBalanceResponse, error) {
	user := &model.Balance{
		BalanceID: uuid.New(), //TODO: return to parsed UserID when Profile is ready
		Balance:   req.Balance.Balance,
	}
	err := h.srv.CreateBalance(ctx, user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user}).Errorf("CreateBalance: %v", err)
		return nil, fmt.Errorf("CreateBalance: %w", err)
	}
	return &proto.CreateBalanceResponse{}, nil
}

// DeleteUserBalance deletes user's balance
func (h *BalanceHandler) DeleteUserBalance(ctx context.Context, req *proto.DeleteBalanceRequest) (*proto.DeleteBalanceResponse, error) {
	err := h.CustomIDValidaion(ctx, req.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ProfileID": req.ProfileID}).Errorf("Validate: %v", err)
		return nil, fmt.Errorf("validate: %w", err)
	}
	ID, err := uuid.Parse(req.ProfileID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ProfileID.ID": req.ProfileID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	err = h.srv.DeleteBalance(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": ID}).Errorf("DeleteBalance: %v", err)
		return nil, fmt.Errorf("DeleteBalance: %w", err)
	}
	return &proto.DeleteBalanceResponse{}, nil
}

// GetAllUserBalances returns all user balances
func (h *BalanceHandler) GetAllUserBalances(ctx context.Context, _ *proto.GetAllBalanceRequest) (*proto.GetAllBalanceResponse, error) {
	users, err := h.srv.GetAllBalances(ctx)
	if err != nil {
		logrus.Errorf("GetAllBalances: %v", err)
		return nil, fmt.Errorf("GetAllBalances: %w", err)
	}
	response := []*proto.Balance{}
	for _, user := range users {
		response = append(response, &proto.Balance{
			ProfileID: user.BalanceID.String(),
			Balance:   user.Balance,
		})
	}
	return &proto.GetAllBalanceResponse{Balances: response}, nil
}
