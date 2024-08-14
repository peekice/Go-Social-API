package gateways

import (
	service "go-api/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	UserService service.IUsersService
	PostService service.IPostsService
}

func NewHTTPGateway(app *fiber.App, users service.IUsersService, posts service.IPostsService) {
	gateway := &HTTPGateway{
		UserService: users,
		PostService: posts,
	}

	GatewayUsers(*gateway, app)
}
