package main

import (
	"j_study_blog/api"
	"j_study_blog/j_dict"
	"j_study_blog/repository"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "import:dicts" {
		repo := repository.NewVocabRepo()
		j_dict.Import(repo)
	} else {
		app := fiber.New()
		api.RegsiterV1Routes(app)

		app.Listen(":3000")
	}
}
