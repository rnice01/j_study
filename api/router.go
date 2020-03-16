package api

import (
	"github.com/gofiber/fiber"
	apiV1 "j_study_blog/api/v1"
	"j_study_blog/repository"
)

func RegsiterV1Routes(app *fiber.App) {
	repo := repository.NewVocabRepo()
	vocabController := apiV1.NewVocabController(&repo)

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/vocabs/kanji/:kanji", vocabController.FindVocab)
}
