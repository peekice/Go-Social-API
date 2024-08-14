package gateways

import (
	"go-api/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func GatewayUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", gateway.Register)
	api.Post("/login", gateway.Login)
	api.Get("/all_posts", gateway.GetAllPosts)

	api.Use(middlewares.SetJWtHeaderHandler())
	api.Post("/create_post", gateway.CreatePost)
	api.Put("/edit_post", gateway.EditPost)
	api.Put("like_post", gateway.LikePost)
	api.Put("/comment_post", gateway.CommentPost)
	api.Delete("/delete_post", gateway.DeletePost)
	api.Put("/edit_comment", gateway.EditComment)
	api.Delete("/delete_comment", gateway.DeleteComment)
}
