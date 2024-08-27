package service

import (
	"tz-kode/internal/entity"
	"tz-kode/internal/repository"
)

type NotesService struct {
	repo repository.Notes
}

func NewNotesService(repo repository.Notes) *NotesService {
	return &NotesService{repo: repo}
}

func (a *NotesService) CreateNote(note *entity.Note) error {
	return a.repo.CreateNote(note)
}
func (a *NotesService) GetAll(userId string) ([]entity.Note, error) {
	return a.repo.GetAll(userId)

}
