package handler

import (
	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/products/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	usecase ports.ProductUsecase
}

func NewHandler(u ports.ProductUsecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) SaveCategory(c *fiber.Ctx) error {
	var req dto.CreateProductsRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	// Validate Request
	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateCategory(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.Id == uuid.Nil {
		return response.SuccessWithDuplicate(c, "category already exist")
	}

	return response.Created(c, result, "save category success")
}

func (h *Handler) SaveBrand(c *fiber.Ctx) error {
	var req dto.CreateProductsRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	// Validate Request
	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateBrand(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.Id == uuid.Nil {
		return response.SuccessWithDuplicate(c, "brand already exist")
	}

	return response.Created(c, result, "save brand success")
}

func (h *Handler) SaveModel(c *fiber.Ctx) error {
	var req dto.CreateProductsRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	// Validate Request
	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateModel(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.Id == uuid.Nil {
		return response.SuccessWithDuplicate(c, "model already exist")
	}

	return response.Created(c, result, "save model success")
}

func (h *Handler) GetCategoryList(c *fiber.Ctx) error {

	result, err := h.usecase.ListCategories()
	if err != nil {
		return response.Error(c, 500, err)
	}

	if len(result) == 0 {
		return response.SuccessWithDuplicate(c, "category already exist")
	}

	return response.Created(c, result, "get category success")
}

func (h *Handler) GetBrandList(c *fiber.Ctx) error {

	result, err := h.usecase.ListBrands()
	if err != nil {
		return response.Error(c, 500, err)
	}

	if len(result) == 0 {
		return response.SuccessWithDuplicate(c, "brand already exist")
	}

	return response.Created(c, result, "get brand success")
}

func (h *Handler) GetModelList(c *fiber.Ctx) error {

	result, err := h.usecase.ListModels()
	if err != nil {
		return response.Error(c, 500, err)
	}

	if len(result) == 0 {
		return response.SuccessWithDuplicate(c, "model already exist")
	}

	return response.Created(c, result, "get model success")
}

func (h *Handler) SaveProduct(c *fiber.Ctx) error {
	var req dto.CreateProducts

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	// Validate Request
	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateProduct(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.Id == uuid.Nil {
		return response.SuccessWithDuplicate(c, "product already exist")
	}

	return response.Created(c, result, "save product success")
}
