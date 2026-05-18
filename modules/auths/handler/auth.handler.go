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
		return response.Error(c, 500, err)
	}

	ctx := c.Context()

	result, err := h.service.Login(ctx, req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	return response.Success(c, result)
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
