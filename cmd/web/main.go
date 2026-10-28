package main

import (
	"belajar-clean-arsitektur/internal/config"
	"belajar-clean-arsitektur/internal/delivery/rest"
	"belajar-clean-arsitektur/internal/delivery/rest/route"
	"belajar-clean-arsitektur/internal/repository"
	"belajar-clean-arsitektur/internal/usecase"
)

func main() {
	cfg := config.NewConfig()
	validate := config.NewValidate()
	db := config.NewDB(cfg.DB)
	log := config.NewLogger(cfg.Log)
	app := config.NewApp(cfg.App)
	hasher := config.NewHasher()

	userRepository := repository.NewUserRepository(log, db)
	userUseCase := usecase.NewUserUseCase(log, validate, hasher, userRepository)
	userController := rest.NewUserController(log, app, userUseCase)

	route := route.NewRoute(app, userController)
	route.Setup()

	if err := app.Listen(cfg.App.Addr); err != nil {
		log.WithError(err).Fatal("failed to start server")
	}
}
