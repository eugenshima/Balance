// Package model of our entity
package model

import (
	"github.com/google/uuid"
)

// User struct represents a user model
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" validate:"required"`
	Balance  float64   `json:"balance" validate:"required,min=-1,max=100000"`
}
