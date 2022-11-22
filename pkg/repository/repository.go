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
	Create(userId uuid.UUID, post listing.Post) (uuid.UUID, error)
	GetAll() ([]listing.Post, error)
	GetById(id uuid.UUID) (listing.Post, error)
}
type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
