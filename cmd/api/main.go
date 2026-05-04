package main

import (
	"github.com/BreakDown-CS/erp-setting-cs/internal/config"
	"github.com/BreakDown-CS/erp-setting-cs/internal/database"
	"github.com/BreakDown-CS/erp-setting-cs/modules/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	db := database.ConnPostgres(cfg)

	users.Wire(app, db)

	app.Listen(":8080")
}
