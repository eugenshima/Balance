// Package model of our entity
package model

import (
	"github.com/google/uuid"
)

// User struct represents a user model
type Balance struct {
	User_ID uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}
