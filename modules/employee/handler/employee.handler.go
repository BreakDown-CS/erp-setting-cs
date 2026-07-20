package handler

import (
	"github.com/google/uuid"

	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service posts.EmployeeUsecase
}

func NewHandler(u posts.EmployeeUsecase) *Handler {
	return &Handler{service: u}
}

func (h *Handler) InsertEmployee(c *fiber.Ctx) error {
	req := new(dto.InsetEmpolyeeRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, nil, 500)
	}

	result, err := h.service.InsertEmployee(c.Context(), *req)
	if err != nil {
		return response.Error(c, err, 500)
	}

	if result.EmployeeId == uuid.Nil {
		return response.Warning(c, "พนักงานคนนี้มีอยู่แล้ว", nil)
	}

	return response.Success(c, result, "ดึงข้อมูลสําเร็จ")
}
