package gateways

import (
	"go-api/domain/entities"
	"go-api/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) GetAllPosts(ctx *fiber.Ctx) error {
	posts, err := h.PostService.GetAllPosts()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: posts})
}

func (h *HTTPGateway) CreatePost(ctx *fiber.Ctx) error {
	decodedToken, err := middlewares.DecodeJWTToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: "Unauthorization Token."})
	}
	userID := decodedToken.UserID

	var payloadData = entities.UserPostModel{}
	if err := ctx.BodyParser(&payloadData); err != nil {
		return ctx.Status(422).JSON(entities.ResponseModel{Message: "Unprocessable Entity"})
	}

	err = h.PostService.CreatePost(userID, payloadData)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) EditPost(ctx *fiber.Ctx) error {
	decodedToken, err := middlewares.DecodeJWTToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: "Unauthorization Token."})
	}
	userID := decodedToken.UserID

	parmas := ctx.Queries()
	postID := parmas["post_id"]

	type EditContent struct {
		Content string `json:"content"`
	}

	var newContent = EditContent{}

	if err := ctx.BodyParser(&newContent); err != nil {
		return ctx.Status(422).JSON(entities.ResponseModel{Message: "Unprocessable Entity"})
	}

	err = h.PostService.EditPost(userID, postID, newContent.Content)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) LikePost(ctx *fiber.Ctx) error {
	_, err := middlewares.DecodeJWTToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: "Unauthorization Token."})
	}

	parmas := ctx.Queries()
	postID := parmas["post_id"]

	err = h.PostService.LikePost(postID)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) CommentPost(ctx *fiber.Ctx) error {
	decodedToken, err := middlewares.DecodeJWTToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: "Unauthorization Token."})
	}
	userID := decodedToken.UserID

	parmas := ctx.Queries()
	postID := parmas["post_id"]

	type Comment struct {
		Content string `json:"content"`
	}

	var newComment = Comment{}

	if err := ctx.BodyParser(&newComment); err != nil {
		return ctx.Status(422).JSON(entities.ResponseModel{Message: "Unprocessable Entity"})
	}

	err = h.PostService.CommentPost(userID, postID, newComment.Content)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) DeletePost(ctx *fiber.Ctx) error {
	decodedToken, err := middlewares.DecodeJWTToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entities.ResponseModel{Message: "Unauthorization Token."})
	}
	userID := decodedToken.UserID

	parmas := ctx.Queries()
	postID := parmas["post_id"]

	err = h.PostService.DeletePost(userID, postID)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}
