package service

import (
	"tz-kode/internal/entity"
	"tz-kode/internal/repository"
)

type User interface {
	Create(user *entity.User) error
	FindById(value string) (*entity.User, error)
	FindByEmain(value string) (*entity.User, error)
}

type Notes interface {
	CreateNote(note *entity.Note) error
	GetAll(userId string) ([]entity.Note, error)
}

type Service struct {
	User
	Notes
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(repository.User),
		Notes: NewNotesService(repository.Notes),
	}
}
