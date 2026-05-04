package users

import (
	"users-api/modules/users/handler"
	"users-api/modules/users/repository"
	"users-api/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Wire(app *fiber.App, db *gorm.DB) {
	repo := repository.NewRepository(db)
	usecase := service.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	UserRouter(app, handler)
}
