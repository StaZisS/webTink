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
	GetAll() ([]listing.PostSend, error)
	GetById(id uuid.UUID) (listing.PostSend, error)
	Delete(idUser, idPost uuid.UUID) error
	Update(idUser, idPost uuid.UUID, input listing.UpdatePostInput) error
}
type Email interface {
	SendEmail(email listing.Email) error
}
type Repository struct {
	Authorization
	Post
	Email
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
		Email:         NewEmailPostgres(db),
	}
}
