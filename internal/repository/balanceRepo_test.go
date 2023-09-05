// Package repository contains repository tests in this case
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
	BalanceID: uuid.New(),
	Balance:   1234.25,
}

var wrongTestEntity = model.Balance{
	BalanceID: uuid.Nil,
	Balance:   1234.75,
}

// TestPgxCreateDeleteBalance function tests create and delete methods
func TestPgxCreateDeleteBalance(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	err = rps.DeleteBalance(context.Background(), testEntity.BalanceID)
	require.NoError(t, err)
}

// TestPgxDeleteNilBalance function tests nil deletion of delete method
func TestPgxDeleteNilBalance(t *testing.T) {
	err := rps.DeleteBalance(context.Background(), wrongTestEntity.BalanceID)
	require.Error(t, err)
}

// TestPgxUpdateBalance function tests update method
func TestPgxUpdateBalance(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	testEntity.Balance = 4321
	err = rps.UpdateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	err = rps.DeleteBalance(context.Background(), testEntity.BalanceID)
	require.NoError(t, err)
}

// TestPgxErrorUpdateBalance function tests error update method
func TestPgxErrorUpdateBalance(t *testing.T) {
	err := rps.UpdateBalance(context.Background(), &testEntity)
	require.Error(t, err)
}

// TestGetBalanceByID function tests get method
func TestGetBalanceByID(t *testing.T) {
	err := rps.CreateBalance(context.Background(), &testEntity)
	require.NoError(t, err)
	testResult, err := rps.GetUserByID(context.Background(), testEntity.BalanceID)
	require.NoError(t, err)
	require.NotNil(t, testResult)
	err = rps.DeleteBalance(context.Background(), testEntity.BalanceID)
	require.NoError(t, err)
}

// TestGetBalanceByWrongID function tests error get method
func TestGetBalanceByWrongID(t *testing.T) {
	testResult, err := rps.GetUserByID(context.Background(), testEntity.BalanceID)
	require.Error(t, err)
	require.Nil(t, testResult)
}

// TestGetAllBalances function tests get all method
func TestGetAllBalances(t *testing.T) {
	testResult, err := rps.GetAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, testResult)
}
