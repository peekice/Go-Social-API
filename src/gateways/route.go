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
	api.Post("/create_posts", gateway.CreatePost)
	api.Put("/edit_posts", gateway.EditPost)
	api.Put("like_posts", gateway.LikePost)
	api.Put("/comment_posts", gateway.CommentPost)
	api.Delete("/delete_posts", gateway.DeletePost)
}
