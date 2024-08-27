package service

import (
	"tz-kode/internal/entity"
	"tz-kode/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(article *entity.User) error {
	return s.repo.Create(article)
}

func (s *UserService) FindById(value string) (*entity.User, error) {
	return s.repo.FindById(value)
}

func (s *UserService) FindByEmain(value string) (*entity.User, error) {
	return s.repo.FindByEmain(value)
}
