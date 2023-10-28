package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/jubelio-chatapp/domain"
	"github.com/timopattikawa/jubelio-chatapp/dto"
	"github.com/timopattikawa/jubelio-chatapp/internal/handler/middleware"
	"log"
	"strconv"
)

type ChatHandler struct {
	service domain.ChatService
}

func NewChatHandler(app *fiber.App, service domain.ChatService) {
	h := ChatHandler{
		service: service,
	}
	app.Post("/message/send", middleware.JWTProtected(), h.sendMessage)
	app.Get("/message/history", middleware.JWTProtected(), h.fetchAllMessageFromUser)
}

func (c ChatHandler) sendMessage(ctx *fiber.Ctx) error {
	var chatDto dto.ChatDto
	if err := ctx.BodyParser(&chatDto); err != nil {
		log.Println("Fail Parse Body Auth Request", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong body request",
		})
	}

	if chatDto.Sender == 0 || chatDto.Receiver == 0 || chatDto.Message == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fail body request",
		})
	}

	message, err := c.service.SendMessage(chatDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(message)
}

func (c ChatHandler) fetchAllMessageFromUser(ctx *fiber.Ctx) error {
	idReceiver := ctx.Query("receiver")
	idSender := ctx.Query("sender")
	if idReceiver == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fail param",
		})
	}

	if idSender == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fail param",
		})
	}

	theIDReceiver, err := strconv.Atoi(idReceiver)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	theIdSender, err := strconv.Atoi(idSender)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, err := c.service.SearchHistoryByReceiver(uint(theIdSender), uint(theIDReceiver))
	res := dto.ChatRes{
		Sender:   theIdSender,
		Receiver: theIDReceiver,
		Messages: message,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
