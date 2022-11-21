package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	listing "web"
)

type Authorization interface {
	CreateUser(user listing.User) (uuid.UUID, error)
	GetUser(email, password string) (listing.User, error)
	GetUserById(id uuid.UUID) (listing.User, error)
}
type Post interface {
}
type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
