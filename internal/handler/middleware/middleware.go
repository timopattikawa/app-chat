package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"log"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte("the secret of app"),
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
