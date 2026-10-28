package rest

import (
	"belajar-clean-arsitektur/internal/entity"
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type RequestUserRegister struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type ResponseUser struct {
	Email string `json:"email"`
}

type UserUseCase interface {
	Create(ctx context.Context, request *entity.User) error
}

type UserController struct {
	Log         *logrus.Logger
	App         *fiber.App
	UserUseCase UserUseCase
}

func NewUserController(log *logrus.Logger, app *fiber.App, userUseCase UserUseCase) *UserController {
	return &UserController{
		Log:         log,
		App:         app,
		UserUseCase: userUseCase,
	}
}

func (c *UserController) Register(ctx fiber.Ctx) error {
	request := new(RequestUserRegister)

	if err := ctx.Bind().Body(&request); err != nil {
		c.Log.WithError(err).Error("invalid request register")
		return err
	}

	var user = new(entity.User)
	user.Email = request.Email
	user.Password = request.Password

	if err := c.UserUseCase.Create(ctx, user); err != nil {
		return err
	}

	var response = new(ResponseUser)
	response.Email = user.Email

	return ctx.JSON(WebResponse[*ResponseUser]{Data: response})
}
