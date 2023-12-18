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
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository()

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Validate, userRepository)

	// setup controller
	userController := http.NewUserController(userUseCase)

	authMiddleware := middleware.JWTProtected()

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
