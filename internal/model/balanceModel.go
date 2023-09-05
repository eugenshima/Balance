// Package model of our entity
package model

import (
	"github.com/google/uuid"
)

// Balance struct represents a user model
type Balance struct {
	BalanceID uuid.UUID `json:"user_id"`
	ProfileID uuid.UUID `json:"profile_id"`
	Balance   float64   `json:"balance"`
}
