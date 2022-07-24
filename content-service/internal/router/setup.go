package router

import "github.com/gofiber/fiber/v2"

func InstallRouter(app *fiber.App) {
	setup(app,
		NewVideoRouter(),
	)
}

func setup(app *fiber.App, router ...Router) {
	api := app.Group("/api/v1")
	for _, r := range router {
		r.InstallRouter(api)
	}
}
