package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func SetupStatic(app *fiber.App) {
	// serve Single Page application on "/web"
	// assume static file at dist folder
	app.Static("/assets", "dist/assets")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("dist/index.html")
	})
}
