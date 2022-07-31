package controllers

import (
	"github.com/giovannitgl/video-services/content-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func UploadVideo(c *fiber.Ctx) error {
	params := c.AllParams()
	videoId := params["id"]
	videoIdInt, err := strconv.ParseUint(videoId, 10, 32)
	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, "invalid id typem should be integer")
	}

	video := service.VideoGetOne(uint(videoIdInt))
	if video == nil {
		return returnErrorMessage(c, fiber.StatusNotFound, "video does not exists")
	}

	file, err := c.FormFile("fileUpload")

	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return returnErrorMessage(c, fiber.StatusBadRequest, err.Error())
	}
	defer buffer.Close()

	info, err := service.UploadFile(file, buffer, err)

	if err != nil {
		return returnErrorMessage(c, fiber.StatusInternalServerError, err.Error())
	}

	err = service.SendVideoUploadedEvent(uint(videoIdInt), "")
	if err != nil {
		log.Printf("Could not send video upload event for id %d, %s", videoIdInt, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": info,
	})
}
