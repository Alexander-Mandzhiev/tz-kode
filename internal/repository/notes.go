package repository

import (
	"fmt"
	"time"
	"tz-kode/internal/entity"

	"github.com/jmoiron/sqlx"
)

type NotesPostgres struct {
	db *sqlx.DB
}

func NewNotesPostgres(db *sqlx.DB) *NotesPostgres {
	return &NotesPostgres{db: db}
}

func (a *NotesPostgres) CreateNote(note *entity.Note) error {
	query := "INSERT INTO notes (id, text, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, text, user_id, created_at;"
	fmt.Print()
	return a.db.QueryRow(query, note.ID, note.Text, note.UserId, time.Now()).
		Scan(&note.ID, &note.Text, &note.UserId, &note.CreatedAt)
}
func (a *NotesPostgres) GetAll(userId string) ([]entity.Note, error) {
	n := []entity.Note{}

	query := `SELECT * FROM notes WHERE user_id = $1`

	err := a.db.Select(&n, query, userId)
	return n, err
}
