package handler

import (
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/setup/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service ports.SetupUsecase
}

func NewHandler(u ports.SetupUsecase) *Handler {
	return &Handler{service: u}
}

func (h *Handler) GetAllBranches(c *fiber.Ctx) error {
	result, err := h.service.GetAllBranches(c.Context())
	if err != nil {
		return response.Error(c, err, 500)
	}

	return response.Success(c, result, "ดึงข้อมูลสําเร็จ")
}

func (h *Handler) GetAllDepartment(c *fiber.Ctx) error {
	result, err := h.service.GetAllDepartment(c.Context())
	if err != nil {
		return response.Error(c, err, 500)
	}

	return response.Success(c, result, "ดึงข้อมูลสําเร็จ")
}

func (h *Handler) GetAllEmployeesStatus(c *fiber.Ctx) error {
	result, err := h.service.GetAllEmployeesStatus(c.Context())
	if err != nil {
		return response.Error(c, err, 500)
	}

	return response.Success(c, result, "ดึงข้อมูลสําเร็จ")
}

func (h *Handler) GetAllPositions(c *fiber.Ctx) error {
	result, err := h.service.GetAllPositions(c.Context())
	if err != nil {
		return response.Error(c, err, 500)
	}

	return response.Success(c, result, "ดึงข้อมูลสําเร็จ")
}
