package listing

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID         uuid.UUID `json:"id" db:"id"`
	AuthorID   uuid.UUID `json:"author_id" db:"author_id"`
	Name       string    `json:"name" db:"name"`
	Surname    string    `json:"surname" db:"surname"`
	Role       string    `json:"role" db:"role"`
	Education  string    `json:"education" db:"education"`
	Additional string    `json:"additional" db:"additional"`
	Photo      string    `json:"photo" db:"photo"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type UpdatePostInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Role       string `json:"role"`
	Education  string `json:"education"`
	Additional string `json:"additional"`
	Photo      string `json:"photo"`
}
