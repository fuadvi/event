package route

import (
	"github.com/gofiber/fiber/v2"
	"rental_mobile_fiber/internal/delivery/http"
)

//func UserRoute(app *fiber.App) {
//	routes := app.Group("users")
//	//routes.Post("/", middleware.JWTProtected(), http.UserController{}.Register)
//	routes.Post("/", http.UserController{}.Register)
//}

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupUserRoute()
}

func (c *RouteConfig) SetupUserRoute() {
	c.App.Post("/api/users", c.UserController.Register)
}
