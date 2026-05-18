package auth

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/handler"
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/repository"
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Wire(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo, db)
	handler := handler.NewHandler(usecase)

	authRouter(app, handler)
}
