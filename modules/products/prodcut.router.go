package products

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/handler"

	"github.com/gofiber/fiber/v2"
)

func productRouter(app *fiber.App, handler *handler.Handler) {
	api := app.Group("/setting/products")

	api.Post("/saveCategory", handler.SaveCategory)
	api.Post("/saveBrand", handler.SaveBrand)
	api.Post("/saveModel", handler.SaveModel)
	api.Get("/getCategory", handler.GetCategoryList)
	api.Get("/getBrand", handler.GetBrandList)
	api.Get("/getModel", handler.GetModelList)
	api.Post("/SaveProduct", handler.SaveProduct)
}
