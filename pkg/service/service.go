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
	GetAll() ([]listing.Post, error)
	GetById(id uuid.UUID) (listing.Post, error)
	Delete(idUser, idPost uuid.UUID) error
	Update(idUser, idPost uuid.UUID, input listing.UpdatePostInput) error
}
type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
	}
}
