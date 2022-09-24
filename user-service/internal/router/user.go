package router

import (
	"github.com/giovannitgl/video-services/user-service/app/v1/controllers"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (r UserRouter) InstallRouter(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/", controllers.RenderUser)
	user.Post("/", controllers.CreateUser)
	user.Post("/auth/login", controllers.LoginUser)
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}
