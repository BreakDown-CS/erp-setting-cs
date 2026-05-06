package handler

import (
	"strconv"

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
		return response.Error(c, 400, err)
	}

	ctx := c.Context()

	result, err := h.service.CreateStaff(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.StaffId == uuid.Nil {
		return response.SuccessWithDuplicate(c, "staff already exist")
	}

	return response.Created(c, result, "save staff success")
}

func (h *Handler) GetStaffList(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "50"))
	if err != nil || limit <= 0 {
		limit = 50
	}
	staffs, total, err := h.service.GetStaffList(page, limit)
	if err != nil {
		return response.Error(c, 500, err)
	}

	meta := &response.Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return response.SuccessWithMeta(c, staffs, meta)
}

// func (h *Handler) GetStaffById(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return response.Error(c, 400, err)
// 	}

// 	staff, err := h.service.GetStaffById(id)
// 	if err != nil {
// 		return response.Error(c, 500, err)
// 	}

// 	return response.Success(c, staff)
// }
