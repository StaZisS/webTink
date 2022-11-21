package listing

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"-"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Grade     int32     `json:"grade"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}