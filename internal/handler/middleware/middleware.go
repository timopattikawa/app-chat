package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
	"log"
)

func JWTProtected() func(*fiber.Ctx) error {
	jwtKey := viper.GetString("jwt_key")
	config := jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte(jwtKey),
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	log.Println("token: ", c.Locals("token"))
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
