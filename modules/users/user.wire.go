package users

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/users/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/users/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Wire(app *fiber.App, db *gorm.DB) {
	repo := repository.NewRepository(db)
	usecase := service.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	UserRouter(app, handler)
}
