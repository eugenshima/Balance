// Package repository contains CRUD operations
package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// PsqlConnection is a struct, which contains Pool variable
type PsqlConnection struct {
	pool *pgxpool.Pool
}

// NewPsqlConnection constructor for PsqlConnection
func NewPsqlConnection(pool *pgxpool.Pool) *PsqlConnection {
	return &PsqlConnection{pool: pool}
}

// GetUserByID function returns user information by ID
func (db *PsqlConnection) GetUserByID(ctx context.Context, profileID uuid.UUID) (*model.Balance, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	var profile model.Balance
	err = tx.QueryRow(ctx, "SELECT profile_id, balance FROM shares.balance WHERE profile_id = $1", profileID).Scan(&profile.BalanceID, &profile.Balance)
	if err != nil || profile.BalanceID == uuid.Nil {
		return nil, fmt.Errorf("QueryRow(): %w", err)
	}
	return &profile, nil
}

// GetAll function executes SQL request to select all rows from Database
func (db *PsqlConnection) GetAll(ctx context.Context) ([]*model.Balance, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	rows, err := tx.Query(ctx, "SELECT balance_id, profile_id, balance FROM shares.balance")
	if err != nil {
		return nil, fmt.Errorf("Query(): %w", err)
	}
	defer rows.Close()

	// Create slice to store data from our SQL request
	var results []*model.Balance

	// go through each line
	for rows.Next() {
		user := &model.Balance{}
		err := rows.Scan(&user.BalanceID, &user.ProfileID, &user.Balance)
		if err != nil {
			return nil, fmt.Errorf("Scan(): %w", err) // Returning error message
		}
		results = append(results, user)
	}
	return results, rows.Err()
}

// UpdateBalance function updates user's balance information
func (db *PsqlConnection) UpdateBalance(ctx context.Context, balance *model.Balance) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	err = tx.QueryRow(ctx, "SELECT balance_id FROM shares.balance WHERE profile_id = $1", balance.ProfileID).Scan(&balance.BalanceID)
	if err != nil || balance.ProfileID == uuid.Nil {
		return fmt.Errorf("QueryRow(): %w", err)
	}
	tag, err := tx.Exec(ctx, "UPDATE shares.balance SET balance = $1 WHERE balance_id = $2", balance.Balance, balance.BalanceID)
	if err != nil || tag.RowsAffected() == 0 {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// CreateBalance function creates user's balance
func (db *PsqlConnection) CreateBalance(ctx context.Context, balance *model.Balance) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	_, err = tx.Exec(ctx, "INSERT INTO shares.balance VALUES ($1, $2, $3)", balance.BalanceID, balance.ProfileID, balance.Balance)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// DeleteBalance function deletes user's balance
func (db *PsqlConnection) DeleteBalance(ctx context.Context, ProfileID uuid.UUID) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	balanceID := uuid.New()
	err = tx.QueryRow(ctx, "SELECT balance_id FROM shares.balance WHERE profile_id = $1", ProfileID).Scan(&balanceID)
	if err != nil || ProfileID == uuid.Nil {
		return fmt.Errorf("QueryRow(): %w", err)
	}
	tag, err := tx.Exec(ctx, "DELETE FROM shares.balance WHERE balance_id = $1", balanceID)
	if err != nil || tag.RowsAffected() == 0 {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
