package auth

import (
	"github.com/BreakDown-CS/erp-setting-cs/middlewares"
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/handler"
	"github.com/gofiber/fiber/v2"
)

func authRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/auth")

	api.Post("/login", handler.Login)
	api.Get("/auth-test", middlewares.JwtAuthentication(), handler.AuthTest)
}
