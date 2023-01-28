package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupMiddlewares(app *fiber.App) {
	// Disallow request from anywhere but packform-test
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://packform-test.mzaidannas.me",
		AllowCredentials: true,
	}))
	// Enable caching of GET routes
	// app.Use(cache.New(cache.Config{
	// 	CacheControl: true, // Enable client-side cache too
	// }))
	// Enables response compression according to the Accept-Encoding header
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
}
