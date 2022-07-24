package controllers

import (
	"github.com/giovannitgl/video-services/content-service/app/v1/schema"
	"github.com/giovannitgl/video-services/content-service/internal/entities"
	"github.com/giovannitgl/video-services/content-service/internal/service"
	"github.com/giovannitgl/video-services/content-service/internal/validator"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetVideos(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	videos := service.VideoPaginatedFilter(pagination)
	videoFormat := []schema.Video{}
	for _, v := range videos {
		videoFormat = append(videoFormat, *schema.VideoResponse(v))
	}

	return c.JSON(videoFormat)
}

func CreateVideo(c *fiber.Ctx) error {
	videoCreate := new(schema.VideoCreate)

	if err := c.BodyParser(videoCreate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	video := entities.Video{
		Title:        videoCreate.Title,
		Description:  videoCreate.Description,
		CreationTime: time.Now(),
	}

	errors := validator.ValidateModel(video)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			errors,
		)
	}

	err := service.VideoCreate(&video)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(schema.VideoResponse(video))
}

func UpdateVideo(c *fiber.Ctx) error {
	videoUpdate := new(schema.VideoUpdate)

	if err := c.BodyParser(videoUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	errors := validator.ValidateModel(videoUpdate)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			errors,
		)
	}

	video := service.VideoGetOne(videoUpdate.ID)
	if video == nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": "Video not found",
			})
	}
	video.Title = videoUpdate.Title
	video.Description = videoUpdate.Description

	err := service.VideoUpdate(video)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(schema.VideoResponse(*video))
}
