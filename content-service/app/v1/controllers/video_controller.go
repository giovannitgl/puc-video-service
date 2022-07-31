package controllers

import (
	"github.com/giovannitgl/video-services/content-service/app/v1/schema"
	"github.com/giovannitgl/video-services/content-service/internal/entities"
	"github.com/giovannitgl/video-services/content-service/internal/service"
	"github.com/giovannitgl/video-services/content-service/internal/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func GetVideos(c *fiber.Ctx) error {
	pagination := GetPagination(c)
	videos := service.VideoPaginatedFilter(pagination)
	var videoFormat []schema.Video
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
		return returnErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(schema.VideoResponse(video))
}

func UpdateVideo(c *fiber.Ctx) error {
	videoUpdate := new(schema.VideoUpdate)

	if err := c.BodyParser(videoUpdate); err != nil {
		return returnErrorMessage(c, fiber.StatusInternalServerError, err.Error())
	}

	errors := validator.ValidateModel(videoUpdate)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			errors,
		)
	}

	video := service.VideoGetOne(videoUpdate.ID)
	if video == nil {
		return returnErrorMessage(c, fiber.StatusNotFound, "video not found")
	}
	video.Title = videoUpdate.Title
	video.Description = videoUpdate.Description

	err := service.VideoUpdate(video)
	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(schema.VideoResponse(*video))
}

func PublishVideo(c *fiber.Ctx) error {
	params := c.AllParams()
	videoId := params["id"]
	videoIdInt, err := strconv.ParseUint(videoId, 10, 32)
	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, "invalid id typem should be integer")
	}

	video := service.VideoGetOne(uint(videoIdInt))
	if video == nil {
		return returnErrorMessage(c, fiber.StatusNotFound, "video not found")
	}

	if video.Published {
		return returnErrorMessage(c, fiber.StatusBadRequest, "video already published")
	}

	if video.VideoUrl == "" {
		return returnErrorMessage(c, fiber.StatusBadRequest, "video pending file upload")
	}

	video.Published = true
	err = service.VideoUpdate(video)
	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(schema.VideoResponse(*video))
}
