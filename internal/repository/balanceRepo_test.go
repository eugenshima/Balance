package repository

import (
	"context"
	"testing"

	"github.com/eugenshima/Balance/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var rps *PsqlConnection

var testEntity = model.Balance{
	User_ID: uuid.New(),
	Balance: 1234.25,
}

var wrongTestEntity = model.Balance{
	User_ID: uuid.Nil,
	Balance: 1234.75,
}

func TestPgxCreateDeleteBalance(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	err = rps.DeleteBalance(context.Background(), testEntity.User_ID)
	require.NoError(t, err)
}

func TestPgxDeleteNilBalance(t *testing.T) {
	err := rps.DeleteBalance(context.Background(), wrongTestEntity.User_ID)
	require.Error(t, err)
}

func TestPgxUpdateBalance(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	testEntity.Balance = 4321
	err = rps.UpdateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	err = rps.DeleteBalance(context.Background(), testEntity.User_ID)
	require.NoError(t, err)
}

func TestPgxErrorUpdateBalance(t *testing.T) {
	err := rps.UpdateBalance(context.Background(), &testEntity)
	require.Error(t, err)
}

func TestGetBalanceByID(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	testResult, err := rps.GetUserByID(context.Background(), testEntity.User_ID)
	require.NoError(t, err)
	require.NotNil(t, testResult)
	err = rps.DeleteBalance(context.Background(), testEntity.User_ID)
	require.NoError(t, err)
}

func TestGetBalanceByWrongID(t *testing.T) {
	testResult, err := rps.GetUserByID(context.Background(), testEntity.User_ID)
	require.Error(t, err)
	require.Nil(t, testResult)
}

func TestGetAllBalances(t *testing.T) {
	testResult, err := rps.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, testResult)
}
