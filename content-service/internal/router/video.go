package router

import (
	"github.com/giovannitgl/video-services/content-service/app/v1/controllers"
	"github.com/gofiber/fiber/v2"
)

type VideoRouter struct{}

func (r VideoRouter) InstallRouter(api fiber.Router) {
	video := api.Group("/video")
	video.Get("/", controllers.GetVideos)
	video.Post("/", controllers.CreateVideo)
	video.Put("/", controllers.UpdateVideo)
}

func NewVideoRouter() *VideoRouter {
	return &VideoRouter{}
}
