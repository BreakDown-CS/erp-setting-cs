package setup

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo, db)
	handler := handler.NewHandler(usecase)

	setupRputer(app, handler)
}
