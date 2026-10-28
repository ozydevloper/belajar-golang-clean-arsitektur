package usecase

import (
	"belajar-clean-arsitektur/internal/entity"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}

type UserUseCase struct {
	Log            *logrus.Logger
	Validate       *validator.Validate
	Hasher         Hasher
	UserRepository UserRepository
}

func NewUserUseCase(log *logrus.Logger, validate *validator.Validate, hasher Hasher, userRepository UserRepository) *UserUseCase {
	return &UserUseCase{
		Log:            log,
		Validate:       validate,
		Hasher:         hasher,
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Create(ctx context.Context, request *entity.User) error {
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("invalid request register")
		return entity.ErrBadRequest
	}

	hashed, err := u.Hasher.HashPassword(request.Password)
	if err != nil {
		u.Log.WithError(err).Error("failed to hashing passowrd")
		return entity.ErrInternalServerError
	}

	var user = new(entity.User)
	user.Email = request.Email
	user.Password = hashed

	if err := u.UserRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

type Hasher interface {
	HashPassword(passowrdString string) (string, error)
}
