package utils

import "github.com/gofiber/fiber/v2"

func JSONSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

// Untuk response error
func JSONError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  false,
		"message": message,
		"data":    nil,
	})
}
