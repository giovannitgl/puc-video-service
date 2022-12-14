package content

import (
	"github.com/giovannitgl/video-services/content-service/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewApplication() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	router.InstallRouter(app)
	return app
}
