package api_v1

import (
	"j_study_blog/dictionary"
	"j_study_blog/repository"
	"j_study_blog/web"

	"github.com/gofiber/fiber"
)

type VocabController struct {
	vocabRepo repository.IVocabRepo
}

func NewVocabController(vocabRepo repository.IVocabRepo) VocabController {
	return VocabController{vocabRepo}
}

func (c *VocabController) FindVocab(ctx *fiber.Ctx) {
	filter := dictionary.Vocab{KanjiReading: ctx.Body("kanji_reading")}
	vocab, err := c.vocabRepo.FindBy(filter)

	if err != nil {
		ctx.JSON(web.NewClientError("no results matching your query for '%v'", ctx.Params("kanji")))
	} else {
		ctx.JSON(vocab)
	}
}

