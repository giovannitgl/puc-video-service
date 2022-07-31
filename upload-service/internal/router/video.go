package router

import (
	"github.com/giovannitgl/video-services/content-service/app/v1/controllers"
	"github.com/gofiber/fiber/v2"
)

type VideoRouter struct{}

func (r VideoRouter) InstallRouter(api fiber.Router) {
	video := api.Group("/video")
	video.Post("/:id/upload", controllers.UploadVideo)
}

func NewUploadVideoRouter() *VideoRouter {
	return &VideoRouter{}
}
