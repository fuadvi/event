package config

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"rental_mobile_fiber/internal/delivery/http"
	"rental_mobile_fiber/internal/delivery/http/middleware"
	"rental_mobile_fiber/internal/delivery/http/route"
	"rental_mobile_fiber/internal/repository"
	"rental_mobile_fiber/internal/usecase"
)

type BootstrapConfig struct {
	DB       *sql.DB
	App      *fiber.App
	Validate *validator.Validate
	Viper    *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository()
	eventRepostory := repository.NewEventRepository()

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Validate, userRepository, config.Viper)
	eventUseCase := usecase.NewEventUseCase(config.DB, config.Validate, eventRepostory, config.Viper)

	// setup controller
	userController := http.NewUserController(userUseCase)
	eventController := http.NewEventController(eventUseCase)

	authMiddleware := middleware.JWTProtected(config.Viper)

	routeConfig := route.RouteConfig{
		App:             config.App,
		UserController:  userController,
		EventController: eventController,
		AuthMiddleware:  authMiddleware,
	}

	routeConfig.Setup()
}
