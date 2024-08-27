package test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"tz-kode/internal/entity"
	"tz-kode/internal/repository"
	"tz-kode/pkg/postgres"

	"github.com/google/uuid"
)

func TestUser(t *testing.T) *entity.User {

	return &entity.User{
		ID:        uuid.New().String(),
		Username:  "user@mail.com",
		Email:     "user@mail.com",
		Password:  "124sd2:5f@aFA2",
		CreatedAt: time.Now(),
	}
}

func TestRepository(t *testing.T, databaseURL string) (*repository.Repository, func(...string)) {
	t.Helper()

	db, err := postgres.NewPostgresDB(databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.NewRepository(db)

	return repo, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE ", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
