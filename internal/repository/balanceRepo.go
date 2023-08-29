// Package repository contains CRUD operations
package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"

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
func (db *PsqlConnection) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	var user model.User
	err = db.pool.QueryRow(ctx, "SELECT id, username, balance FROM shares.user WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return nil, fmt.Errorf("QueryRow(): %w", err)
	}
	return &user, nil
}

// GetAll function executes SQL request to select all rows from Database
func (db *PsqlConnection) GetAll(ctx context.Context) ([]*model.User, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	rows, err := db.pool.Query(ctx, "SELECT id, username, balance FROM shares.user")
	if err != nil {
		return nil, fmt.Errorf("Query(): %w", err)
	}
	defer rows.Close()

	// Create slice to store data from our SQL request
	var results []*model.User

	// go through each line
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Balance)
		if err != nil {
			return nil, fmt.Errorf("Scan(): %w", err) // Returning error message
		}
		results = append(results, user)
	}
	return results, rows.Err()
}

// UpdateBalance function updates user's balance information
func (db *PsqlConnection) UpdateBalance(ctx context.Context, user *model.User) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	_, err = db.pool.Exec(ctx, "UPDATE shares.user SET balance = $1, username = $2 WHERE id = $3", user.Balance, user.Username, user.ID)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// CreateBalance function creates user's balance
func (db *PsqlConnection) CreateBalance(ctx context.Context, user *model.User) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	_, err = db.pool.Exec(ctx, "INSERT INTO shares.user VALUES ($1, $2, $3)", user.ID, user.Username, user.Balance)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// DeleteBalance function deletes user's balance
func (db *PsqlConnection) DeleteBalance(ctx context.Context, userID string) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	_, err = db.pool.Exec(ctx, "DELETE FROM shares.user WHERE id = $1", userID)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
