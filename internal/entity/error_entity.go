package entity

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("item already exists")
	ErrBadRequest          = errors.New("invalid request body")
)
