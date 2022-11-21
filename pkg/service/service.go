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
}
type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
