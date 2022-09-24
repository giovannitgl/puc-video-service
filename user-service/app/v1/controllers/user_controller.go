package controllers

import (
	"github.com/giovannitgl/video-services/user-service/app/v1/presenter"
	"github.com/giovannitgl/video-services/user-service/internal/entities"
	userMngr "github.com/giovannitgl/video-services/user-service/internal/manager"
	"github.com/giovannitgl/video-services/user-service/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func RenderUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"test": "ok"})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}
	errors := validator.ValidateUser(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			errors,
		)
	}

	userMngr.RegisterNewUser(user)

	return c.JSON(presenter.UserResponse(*user))
}

func LoginUser(c *fiber.Ctx) error {
	login := new(presenter.Login)
	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	err := userMngr.LoginUser(login.Email, login.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "logged in",
		},
	)

}
