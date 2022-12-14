package service

import (
	"github.com/google/uuid"
	listing "web"
	"web/pkg/repository"
)

type Authorization interface {
	CreateUser(user listing.User) (uuid.UUID, error)
	GenerateToken(email, password string) (string, string, error)
	Refresh(cookie string) (string, error)
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
type Service struct {
	Authorization
	Post
	Email
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
		Email:         NewEmailService(repos.Email),
	}
}
