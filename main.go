package main

import (
	"j_study_blog/api"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	api.RegsiterV1Routes(app)

	app.Listen(8000)
}
