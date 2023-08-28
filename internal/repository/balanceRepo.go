package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/Balance/internal/model"

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

// GetAll function executes SQL request to select all rows from Database
func (db *PsqlConnection) GetAll(ctx context.Context) ([]*model.User, error) {
	rows, err := db.pool.Query(ctx, "SELECT id, username, balance FROM Shares.user")
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
