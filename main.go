package main

import (
	"go-api/configuration"
	ds "go-api/domain/datasources"
	repo "go-api/domain/repositories"
	gw "go-api/src/gateways"
	"go-api/src/middlewares"
	sv "go-api/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)

	userRepo := repo.NewUsersRepository(mongodb)
	postRepo := repo.NewPostsRepository(mongodb)

	sv0 := sv.NewUsersService(userRepo)
	sv1 := sv.NewPostsService(postRepo, userRepo)

	gw.NewHTTPGateway(app, sv0, sv1)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
