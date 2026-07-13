package handler

import (
	"github.com/BreakDown-CS/erp-setting-cs/modules/auths/dto"
	ports "github.com/BreakDown-CS/erp-setting-cs/modules/auths/posts"
	"github.com/BreakDown-CS/erp-setting-cs/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service ports.AuthUsecase
}

func NewHandler(u ports.AuthUsecase) *Handler {
	return &Handler{service: u}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, nil, 500)
	}

	ctx := c.Context()

	result, err := h.service.Login(ctx, req)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return response.Warning(c, "ไม่พบข้อมูลพนักงาน", nil)
		} else if err.Error() == "Unauthorized" {
			return response.Warning(c, "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง", nil)
		} else {
			return response.Error(c, err, 500)
		}
	}

	return response.Success(c, result, "เข้าสู่ระบบสําเร็จ")
}

func (h *Handler) AuthTest(c *fiber.Ctx) error {
	id := c.Locals("user_id")
	username := c.Locals("username")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result": fiber.Map{
			"id":       id,
			"username": username,
		},
	})
}
