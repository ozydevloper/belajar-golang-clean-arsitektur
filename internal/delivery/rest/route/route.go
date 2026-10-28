package route

import (
	"belajar-clean-arsitektur/internal/delivery/rest"

	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *rest.UserController
}

func NewRoute(app *fiber.App, userController *rest.UserController) *RouteConfig {
	return &RouteConfig{
		App:            app,
		UserController: userController,
	}
}

func (r *RouteConfig) SetupGuestRoute() {
	r.App.Post("/api/users/register", r.UserController.Register)
}

func (r *RouteConfig) Setup() {
	r.SetupGuestRoute()
}
