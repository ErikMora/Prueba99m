package main

import (
	"github.com/99minutos/intr-tech-test-a/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := "8080"

	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
