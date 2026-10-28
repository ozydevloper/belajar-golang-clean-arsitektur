package entity

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" validate:"required,max=100"`
	Password  string    `json:"password" validate:"required,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
