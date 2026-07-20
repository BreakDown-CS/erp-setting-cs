package employee

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	EmployeeRouter(app, handler)
}
