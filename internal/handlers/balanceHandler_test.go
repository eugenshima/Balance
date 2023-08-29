// Package handlers contains handler tests in this case
package handlers

import (
	"context"
	"os"
	"testing"

	"github.com/eugenshima/Balance/internal/handlers/mocks"
	"github.com/eugenshima/Balance/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	mockBalanceService *mocks.BalanceService
	mockBalanceEntity  = model.Balance{
		UserID:  uuid.New(),
		Balance: 1234.25,
	}
)

// TestMain execute all tests
func TestMain(m *testing.M) {
	mockBalanceService = new(mocks.BalanceService)
	exitVal := m.Run()
	os.Exit(exitVal)
}

// TestCreate is a mocktest for Create method of interface BalanceService
func TestCreateUserBalance(t *testing.T) {
	mockBalanceService.On("CreateBalance", mock.Anything, mock.AnythingOfType("*model.Balance")).Return(nil).Once()

	err := mockBalanceService.CreateBalance(context.Background(), &mockBalanceEntity)
	require.NoError(t, err)

	assertion := mockBalanceService.AssertExpectations(t)
	require.True(t, assertion)
}

// TestDelete is a mocktest for Delete method of interface BalanceService
func TestDelete(t *testing.T) {
	mockBalanceService.On("DeleteBalance", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil).Once()

	err := mockBalanceService.DeleteBalance(context.Background(), mockBalanceEntity.UserID)
	require.NoError(t, err)

	assertion := mockBalanceService.AssertExpectations(t)
	require.True(t, assertion)
}

// TestGetAll is a mocktest for Get All method of interface BalanceService
func TestGetAll(t *testing.T) {
	mockBalanceService.On("GetAllBalances", mock.Anything).Return([]*model.Balance{}, nil).Twice()
	handler := NewBalancehandler(mockBalanceService, nil)

	res, err := mockBalanceService.GetAllBalances(context.Background())
	require.NoError(t, err)
	require.NotNil(t, res)

	results, err := handler.srv.GetAllBalances(context.Background())
	require.NoError(t, err)
	require.NotNil(t, results)
	require.Equal(t, len(res), len(results))

	assertion := mockBalanceService.AssertExpectations(t)
	require.True(t, assertion)
}

// TestUpdate is a mocktest for Update method of interface BalanceService
func TestUpdate(t *testing.T) {
	mockBalanceService.On("UpdateBalance", mock.Anything, mock.AnythingOfType("*model.Balance")).Return(nil).Once()

	err := mockBalanceService.UpdateBalance(context.Background(), &mockBalanceEntity)
	require.NoError(t, err)

	assertion := mockBalanceService.AssertExpectations(t)
	require.True(t, assertion)
}

// TestGetByID is a mocktest for Get By ID method of interface BalanceService
func TestGetByID(t *testing.T) {
	mockBalanceService.On("GetUserByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&model.Balance{}, nil).Once()

	result, err := mockBalanceService.GetUserByID(context.Background(), mockBalanceEntity.UserID)
	require.NoError(t, err)
	require.NotNil(t, result)

	assertion := mockBalanceService.AssertExpectations(t)
	require.True(t, assertion)
}
