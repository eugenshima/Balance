// Package model of our entity
package model

import (
	"github.com/google/uuid"
)

// Balance struct represents a user model
type Balance struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}
