package repository

import (
	"errors"
	"time"
	"tz-kode/internal/entity"
	"tz-kode/pkg/validation"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (a *UserPostgres) Create(user *entity.User) error {

	if err := user.Valiedate(); err != nil {
		return errors.New("invalid email")
	}
	if !validation.IsValid(user.Password) {
		return errors.New("the password must be longer than 7 characters and contain lowercase letters, uppercase letters, numbers and special characters")
	}

	query := "INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, username, email, created_at;"
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return err
	}
	return a.db.QueryRow(query, user.ID, user.Username, user.Email, hash, time.Now()).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

}

func (a *UserPostgres) FindById(id string) (*entity.User, error) {
	u := &entity.User{}
	query := `SELECT * FROM users WHERE id = $1`
	if err := a.db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (a *UserPostgres) FindByEmain(email string) (*entity.User, error) {
	u := &entity.User{}
	query := `SELECT * FROM users WHERE email = $1`
	if err := a.db.QueryRow(query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		return nil, errors.New("not found")
	}
	return u, nil
}

func generatePasswordHash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
