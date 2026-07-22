package handler

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/dto"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	service posts.EmployeeUsecase
}

func NewHandler(u posts.EmployeeUsecase) *Handler {
	return &Handler{service: u}
}

func (h *Handler) InsertEmployee(c *fiber.Ctx) error {
	req := new(dto.InsertEmpolyeeRequest)
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

	return response.Success(c, result, "บันทึกข้อมูลเรียบร้อยแล้ว")
}

func (h *Handler) GetEmployeeList(c *fiber.Ctx) error {
	req := new(dto.GetListEmpolyeeRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, err, 500)
	}

	result, err := h.service.GetEmployeeList(c.Context(), *req)
	if err != nil {
		return response.Error(c, err, 500)
	}

	if len(result.Employees) == 0 {
		return response.Warning(c, "ไม่พบข้อมูล", nil)
	}

	meta := &response.Meta{
		Page:       result.Page,
		Limit:      result.Limit,
		Total:      int(result.TotalEmployees),
		TotalPages: result.TotalPages,
	}

	return response.SuccessList(c, result.Employees, "ค้นหาข้อมูลสำเร็จ", meta)
}

func (h *Handler) GetEmployeeDetailByUUID(c *fiber.Ctx) error {
	userUUID := c.Params("user_uuid")

	result, err := h.service.GetEmployeeDetail(c.Context(), uuid.Must(uuid.Parse(userUUID)))
	if err != nil {
		return response.Error(c, err, 500)
	}

	if result.EmployeeId == uuid.Nil {
		return response.Warning(c, "ไม่พบข้อมูล", nil)
	}

	return response.Success(c, result, "ค้นหาข้อมูลสำเร็จ")
}
