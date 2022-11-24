package listing

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"email,required"`
	Grade     int       `json:"grade" binding:"required"`
	Password  string    `json:"password" binding:"min=6,max=12,required"`
	Role      string    `json:"role"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}
