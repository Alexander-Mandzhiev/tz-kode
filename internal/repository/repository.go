package repository

import (
	"tz-kode/internal/entity"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user *entity.User) error
	FindById(value string) (*entity.User, error)
	FindByEmain(value string) (*entity.User, error)
	//Update(user entity.User) (entity.User, error)
	//Delete(id string) erro
}

type Notes interface {
	CreateNote(note *entity.Note) error
	GetAll(userId string) ([]entity.Note, error)
}

type Repository struct {
	User
	Notes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Notes: NewNotesPostgres(db),
	}
}
