package handler

import (
	"strconv"

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

// SaveStaff handles the creation of a new staff member
// @Summary Create a new staff
// @Description Create a new staff with the provided details
// @Tags staffs
// @Accept json
// @Produce json
// @Param staff body dto.CreateStaffRequest true "Staff details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /staffs [post]
func (h *Handler) SaveStaff(c *fiber.Ctx) error {
	var req dto.CreateStaffRequest

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

	result, err := h.service.CreateStaff(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	if result.StaffId == uuid.Nil {
		return response.SuccessWithDuplicate(c, "staff already exist")
	}

	return response.Created(c, result, "save staff success")
}

// GetStaffList returns a paginated list of staff members
// @Summary Get staff list
// @Tags staffs
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(50)
// @Success 200 {object} response.Response
// @Router /staffs [get]
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

func (h *Handler) GetStaffById(c *fiber.Ctx) error {
	var req dto.GetStaffById

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	staff, err := h.service.GetStaffById(req.ID)
	if err != nil {
		return response.Error(c, 500, err)
	}

	return response.Success(c, staff)
}
