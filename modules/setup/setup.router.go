package setup

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRputer(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/setup")

	api.Get("/get-dropdown-branches", handler.GetAllBranches)
	api.Get("/get-dropdown-department", handler.GetAllDepartment)
	api.Get("/get-dropdown-employees-status", handler.GetAllEmployeesStatus)
	api.Get("/get-dropdown-positions", handler.GetAllPositions)
}
