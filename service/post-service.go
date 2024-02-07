package service

import (
	"errors"
	"math/rand"

	"github.com/REST-API/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {
	repo PostRepository
}

func NewPostService(repository PostRepository) *service {

	return &service{
		repo: repository,
	}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return s.repo.Save(post)

}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
