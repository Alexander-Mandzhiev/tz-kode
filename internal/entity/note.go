package entity

import "time"

type Note struct {
	ID        string    `json:"id,omitempty" db:"id"`
	Text      string    `json:"text" db:"text"`
	UserId    string    `json:"user_id,omitempty" db:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
