package handler

import (
	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service ports.StaffService
}

func NewHandler(u ports.StaffService) *Handler {
	return &Handler{service: u}
}

func (h *Handler) SaveStaff(c *fiber.Ctx) error {
	var req dto.CreateStaffRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Warning(c, "กรุณากรอกข้อมูลให้ครบ", err)
	}

	// Validate Request
	if errors := helper.ValidateStruct(req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	ctx := c.Context()

	result, err := h.service.CreateStaff(ctx, req)
	if err != nil {
		return response.Error(c, err, 500)
	}

	if result.StaffId == uuid.Nil {
		return response.Warning(c, "ชื่อพนักงานซ้ํา", nil)
	}

	return response.Success(c, result, "บันทึกข้อมูลพนักงานสําเร็จ")
}

func (h *Handler) GetStaffList(c *fiber.Ctx) error {
	var req dto.GetStaffListRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Warning(c, "กรุณากรอกข้อมูลให้ครบ", err)
	}

	staffs, total, err := h.service.GetStaffList(req)
	if err != nil {
		return response.Error(c, err, 500)
	}

	meta := &response.Meta{
		Page:       req.Page,
		Limit:      req.Limit,
		Total:      total,
		TotalPages: (total + req.Limit - 1) / req.Limit,
	}

	return response.SuccessList(c, staffs, "ดึงข้อมูลพนักงานสําเร็จ", meta)
}

func (h *Handler) GetStaffById(c *fiber.Ctx) error {
	var req dto.GetStaffById

	if err := c.BodyParser(&req); err != nil {
		return response.Warning(c, "กรุณากรองข้อมูลให้ครบ", nil)
	}

	result, err := h.service.GetStaffById(req.ID)
	if err != nil {
		return response.Error(c, err, 500)
	}

	return response.Success(c, result, "ดึงข้อมูลพนักงานสําเร็จ")
}
