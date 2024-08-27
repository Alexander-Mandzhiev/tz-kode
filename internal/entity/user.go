package entity

import (
	"time"

	validate "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password," db:"password_hash"`
}

type User struct {
	ID        string    `json:"id,omitempty" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password," db:"password_hash"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type UserResponse struct {
	ID        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Valiedate() error {
	return validate.ValidateStruct(u, validate.Field(&u.Email, validate.Required, is.Email))
}

func (u *User) ClearPassword() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
