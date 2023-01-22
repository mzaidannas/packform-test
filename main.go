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
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	// Enable DB and create migrations before forking to multiple threads
	database.ConnectDB()

	// Use faster JSON library
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	middlewares.SetupMiddlewares(app)

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
