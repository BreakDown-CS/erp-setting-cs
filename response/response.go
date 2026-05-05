package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{
		Success: true,
		Data:    data,
	})
}

func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(201).JSON(Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func Error(c *fiber.Ctx, status int, err error) error {
	msg := ""
	if err != nil {
		msg = err.Error()
	}

	return c.Status(status).JSON(Response{
		Success: false,
		Error:   msg,
	})
}

func SuccessWithDuplicate(c *fiber.Ctx, message string) error {
	return c.JSON(Response{
		Success: true,
		Message: message,
	})
}

func SuccessWithMeta(c *fiber.Ctx, data interface{}, meta *Meta) error {
	return c.JSON(Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}
