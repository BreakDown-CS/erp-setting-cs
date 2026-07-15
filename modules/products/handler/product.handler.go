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
		return response.Warning(c, "กรุณากรอกข้อมูลให้ครบ", err)
	}

	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateCategory(ctx, req)
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if result.Id == uuid.Nil {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.Success(c, result, "บันทึกข้อมูลสําเร็จ")
}

func (h *Handler) SaveBrand(c *fiber.Ctx) error {
	var req dto.CreateProductsRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, nil, 500)
	}

	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateBrand(ctx, req)
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if result.Id == uuid.Nil {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.Success(c, result, "บันทึกข้อมูลสําเร็จ")
}

func (h *Handler) SaveModel(c *fiber.Ctx) error {
	var req dto.CreateProductsRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, nil, 500)
	}

	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateModel(ctx, req)
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if result.Id == uuid.Nil {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.Success(c, result, "บันทึกข้อมูลสําเร็จ")
}

func (h *Handler) GetCategoryList(c *fiber.Ctx) error {

	result, err := h.usecase.ListCategories()
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if len(result) == 0 {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.SuccessList(c, result, "ค้นหาข้อมูลสําเร็จ", nil)
}

func (h *Handler) GetBrandList(c *fiber.Ctx) error {

	result, err := h.usecase.ListBrands()
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if len(result) == 0 {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.SuccessList(c, result, "ค้นหาข้อมูลสําเร็จ", nil)
}

func (h *Handler) GetModelList(c *fiber.Ctx) error {

	result, err := h.usecase.ListModels()
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if len(result) == 0 {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.SuccessList(c, result, "ค้นหาข้อมูลสําเร็จ", nil)
}

func (h *Handler) SaveProduct(c *fiber.Ctx) error {
	var req dto.CreateProducts

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, nil, 500)
	}

	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.usecase.CreateProduct(ctx, req)
	if err != nil {
		return response.Error(c, nil, 500)
	}

	if result.Id == uuid.Nil {
		return response.Success(c, nil, "มีข้อมูลอยู่แล้ว")
	}

	return response.Success(c, result, "บันทึกข้อมูลสำเร็จ")
}
