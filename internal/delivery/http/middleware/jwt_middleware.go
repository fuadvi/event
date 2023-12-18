package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func JWTProtected() func(*fiber.Ctx) error {
	configViper := viper.New()
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: configViper.GetString("jwt.secret")},
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
