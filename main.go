package main

import (
	"log"
	"packform-test/src/database"
	"packform-test/src/middlewares"
	"packform-test/src/router"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if running locally
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	// Enable DB and create migrations before forking to multiple threads
	database.ConnectDB()

	// Use faster JSON library
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	middlewares.SetupMiddlewares(app)
	middlewares.SetupStatic(app)

	router.SetupRoutes(app)

	// Enable request body streaming for CSV imports
	// Note that this may have consequences if typical requests are larger
	// than the configured limit as it's uncertain if Fiber will gracefully
	// handle the streaming if it expects a complete body.
	app.Server().StreamRequestBody = true

	log.Fatal(app.Listen(":3000"))
}
