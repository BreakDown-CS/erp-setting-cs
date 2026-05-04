package handler

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/staffs/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service ports.StaffService
}

func NewHandler(u ports.StaffService) *Handler {
	return &Handler{service: u}
}

func (h *Handler) InsetStaff(c *fiber.Ctx) error {
	var req dto.CreateStaffRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	err := h.service.InsertStaff(req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	return response.Created(c, "created")
}

// func (h *Handler) GetStaffs(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit, _ := strconv.Atoi(c.Query("limit", "50"))

// 	staffs, total, err := h.staffcase.GetAllStaffs(page, limit)
// 	if err != nil {
// 		return response.Error(c, 500, err)
// 	}

// 	meta := &response.Meta{
// 		Page:       page,
// 		Limit:      limit,
// 		Total:      total,
// 		TotalPages: (total + limit - 1) / limit,
// 	}

// 	return response.SuccessWithMeta(c, staffs, meta)
// }

// func (h *Handler) GetStaffById(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return response.Error(c, 400, err)
// 	}

// 	staff, err := h.staffcase.GetStaffById(id)
// 	if err != nil {
// 		return response.Error(c, 500, err)
// 	}

// 	return response.Success(c, staff)
// }
