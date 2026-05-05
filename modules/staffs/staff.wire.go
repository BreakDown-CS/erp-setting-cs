package staffs

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/service"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	usecase := service.NewUsecase(repo, db)
	handler := handler.NewHandler(usecase)

	staffRouter(app, handler)
}
