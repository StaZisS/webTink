package service

import (
	"github.com/google/uuid"
	listing "web"
	"web/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(userId uuid.UUID, post listing.Post) (uuid.UUID, error) {
	return s.repo.Create(userId, post)
}

func (s *PostService) GetAll() ([]listing.PostSend, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetById(id uuid.UUID) (listing.PostSend, error) {
	return s.repo.GetById(id)
}

func (s *PostService) Delete(idUser, idPost uuid.UUID) error {
	return s.repo.Delete(idUser, idPost)
}

func (s *PostService) Update(idUser, idPost uuid.UUID, input listing.UpdatePostInput) error {
	return s.repo.Update(idUser, idPost, input)
}
