package products

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/usecase"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo, db)
	handler := handler.NewHandler(usecase)

	productRouter(app, handler)
}
