package staffs

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/handler"

	"github.com/gofiber/fiber/v2"
)

func staffRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/setting/staffs")

	api.Post("/saveStaff", handler.SaveStaff)
	api.Get("/getStaffList", handler.GetStaffList)
	api.Post("/getStaffById", handler.GetStaffById)
}
