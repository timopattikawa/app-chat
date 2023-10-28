package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"log"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(app *fiber.App, service domain.UserService) {
	h := UserHandler{
		service: service,
	}
	app.Post("/registration", h.RegisUser)
	app.Post("/auth", h.AuthUser)
}

func (u UserHandler) RegisUser(c *fiber.Ctx) error {
	var authRequest dto.AuthReq
	if err := c.BodyParser(&authRequest); err != nil {
		log.Println("Fail Parse Body Auth Request", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong body request",
		})
	}

	if authRequest.Username == "" || authRequest.Password == "" {
		log.Println("Null username or password")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong body request",
		})
	}
	res, err := u.service.RegistrationUser(authRequest)
	if err != nil {
		log.Println("Fail to got response from service", err)
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (u UserHandler) AuthUser(c *fiber.Ctx) error {
	var authRequest dto.AuthReq

	if err := c.BodyParser(&authRequest); err != nil {
		log.Println("Fail Parse Body Auth Request", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong body request",
		})
	}
	user, err := u.service.AuthUser(authRequest)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
