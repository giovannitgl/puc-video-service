package controllers

import (
	"github.com/giovannitgl/video-services/content-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetPagination(c *fiber.Ctx) service.Paginator {
	params := c.AllParams()
	page := 0
	pageSize := 0
	if paramPage, ok := params["page"]; ok {
		paramPageInt, err := strconv.Atoi(paramPage)
		if err == nil {
			page = paramPageInt
		}
	}
	if paramPageSize, ok := params["page"]; ok {
		paramPageSizeInt, err := strconv.Atoi(paramPageSize)
		if err == nil {
			page = paramPageSizeInt
		}
	}

	return service.Paginator{Page: page, PageSize: pageSize}
}

func returnErrorMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(
		fiber.Map{
			"message": message,
		},
	)
}
