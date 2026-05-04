package users

import (
	"users-api/modules/users/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/users")

	api.Post("/", handler.InsetUser)
	api.Get("/", handler.GetUsers)
	api.Get("/:id", handler.GetUserById)

}
