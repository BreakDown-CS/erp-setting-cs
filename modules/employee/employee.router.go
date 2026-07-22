package employee

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/handler"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/employee")

	api.Post("/save", handler.InsertEmployee)
	api.Post("/list", handler.GetEmployeeList)
	api.Get("/get-staff-detail/:user_uuid", handler.GetEmployeeDetailByUUID)
}
