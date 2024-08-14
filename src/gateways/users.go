package gateways

import (
	"go-api/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) Register(ctx *fiber.Ctx) error {

	var payloadData = entities.UserRegisterModel{}
	if err := ctx.BodyParser(&payloadData); err != nil {
		return ctx.Status(422).JSON(entities.ResponseModel{Message: "Unprocessable Entity"})
	}

	err := h.UserService.InsertUser(payloadData)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) Login(ctx *fiber.Ctx) error {

	var payloadData = entities.UserLoginModel{}
	if err := ctx.BodyParser(&payloadData); err != nil {
		return ctx.Status(422).JSON(entities.ResponseModel{Message: "Unprocessable Entity"})
	}

	token, err := h.UserService.Login(payloadData)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(token)
}
