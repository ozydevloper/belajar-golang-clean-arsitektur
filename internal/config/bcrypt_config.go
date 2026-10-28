package config

import (
	"belajar-clean-arsitektur/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type HasherConfig struct{}

func NewHasher() *HasherConfig {
	return &HasherConfig{}
}

func (c *HasherConfig) HashPassword(passwordString string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		return "", entity.ErrInternalServerError
	}

	return string(password), nil
}
