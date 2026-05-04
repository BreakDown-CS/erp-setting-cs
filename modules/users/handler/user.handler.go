package handler

import (
	"strconv"
	"users-api/modules/users/dto"
	"users-api/modules/users/service"
	"users-api/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usercase service.Usecase
}

func NewHandler(u service.Usecase) *Handler {
	return &Handler{usercase: u}
}

func (h *Handler) InsetUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, err)
	}

	err := h.usercase.InsetUser(req)
	if err != nil {
		return response.Error(c, 500, err)
	}

	return response.Created(c, "created")
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "50"))

	users, total, err := h.usercase.GetAllUsers(page, limit)
	if err != nil {
		return response.Error(c, 500, err)
	}

	meta := &response.Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return response.SuccessWithMeta(c, users, meta)
}

func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, 400, err)
	}

	user, err := h.usercase.GetUserById(id)
	if err != nil {
		return response.Error(c, 500, err)
	}

	return response.Success(c, user)
}
