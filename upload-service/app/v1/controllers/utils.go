package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func returnErrorMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(
		fiber.Map{
			"message": message,
		},
	)
}
