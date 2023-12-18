package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"rental_mobile_fiber/internal/config"
)

func main() {
	configViper := config.NewViper()
	app := config.NewFiber(configViper)
	validate := config.NewValidator(configViper)
	db := config.NewDB()

	//// Route with JWT middleware
	//app.Get("/tes/:name", middleware.JWTProtected(), func(c *fiber.Ctx) error {
	//	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	//	return c.SendString(msg)
	//})
	//
	//// Route without JWT middleware
	//app.Get("/:name", func(c *fiber.Ctx) error {
	//	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	//	return c.SendString(msg)
	//})
	//route.UserRoute(app)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validate,
		Viper:    configViper,
	})

	webPort := configViper.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
