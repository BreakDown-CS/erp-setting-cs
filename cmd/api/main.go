package main

import (
	"users-api/internal/config"
	"users-api/internal/database"
	"users-api/modules/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	db := database.ConnPostgres(cfg)

	users.Wire(app, db)

	app.Listen(":8080")
}
