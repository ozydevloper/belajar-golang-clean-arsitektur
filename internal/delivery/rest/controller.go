package rest

import (
	"belajar-clean-arsitektur/internal/entity"
	"errors"

	"github.com/gofiber/fiber/v3"
)

type WebResponse[T any] struct {
	Data T `json:"data"`
}

func ErrorHandler(err error) *fiber.Error {
	switch {
	case errors.Is(err, entity.ErrBadRequest):
		return fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Error())
	case errors.Is(err, entity.ErrConflict):
		return fiber.NewError(fiber.StatusConflict, fiber.ErrConflict.Error())
	default:
		return fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Error())
	}
}
