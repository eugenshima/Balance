// Package repository contains CRUD operations
package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
func (db *PsqlConnection) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.Balance, error) {
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
	var user model.Balance
	err = db.pool.QueryRow(ctx, "SELECT user_id, balance FROM shares.balance WHERE user_id = $1", userID).Scan(&user.User_ID, &user.Balance)
	if err != nil || user.User_ID == uuid.Nil {
		return nil, fmt.Errorf("QueryRow(): %w", err)
	}
	return &user, nil
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
	rows, err := db.pool.Query(ctx, "SELECT user_id, balance FROM shares.balance")
	if err != nil {
		return nil, fmt.Errorf("Query(): %w", err)
	}
	defer rows.Close()

	// Create slice to store data from our SQL request
	var results []*model.Balance

	// go through each line
	for rows.Next() {
		user := &model.Balance{}
		err := rows.Scan(&user.User_ID, &user.Balance)
		if err != nil {
			return nil, fmt.Errorf("Scan(): %w", err) // Returning error message
		}
		results = append(results, user)
	}
	return results, rows.Err()
}

// UpdateBalance function updates user's balance information
func (db *PsqlConnection) UpdateBalance(ctx context.Context, user *model.Balance) error {
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
	tag, err := db.pool.Exec(ctx, "UPDATE shares.balance SET balance = $1 WHERE user_id = $2", user.Balance, user.User_ID)
	if err != nil || tag.RowsAffected() == 0 {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// CreateBalance function creates user's balance
func (db *PsqlConnection) CreateBalance(ctx context.Context, user *model.Balance) error {
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
	_, err = db.pool.Exec(ctx, "INSERT INTO shares.balance VALUES ($1, $2)", user.User_ID, user.Balance)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// DeleteBalance function deletes user's balance
func (db *PsqlConnection) DeleteBalance(ctx context.Context, userID uuid.UUID) error {
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
	tag, err := db.pool.Exec(ctx, "DELETE FROM shares.balance WHERE user_id = $1", userID)
	if err != nil || tag.RowsAffected() == 0 {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
