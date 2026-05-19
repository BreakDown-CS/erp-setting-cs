package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ApiResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Result    interface{} `json:"result"`
	Meta      *Meta       `json:"meta,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:    true,
		Message:   message,
		Code:      fiber.StatusOK,
		Result:    data,
		Timestamp: time.Now().Unix(),
	})
}

func SuccessList(c *fiber.Ctx, data interface{}, message string, meta *Meta) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:    true,
		Message:   message,
		Code:      fiber.StatusOK,
		Result:    data,
		Meta:      meta,
		Timestamp: time.Now().Unix(),
	})
}

func Error(c *fiber.Ctx, data interface{}, code int) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:    true,
		Message:   "error",
		Code:      code,
		Result:    data,
		Timestamp: time.Now().Unix(),
	})
}

func Warning(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:    false,
		Message:   message,
		Code:      fiber.StatusOK,
		Result:    data,
		Timestamp: time.Now().Unix(),
	})
}
