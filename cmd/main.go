package main

import (
	"github.com/erfanva/go-hands-on/database"
	"github.com/erfanva/go-hands-on/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	handlers.SetupRoutes(app)

	app.Listen(":3000")
}