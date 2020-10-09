package api_v1

import (
	"j_study_blog/dictionary"
	"j_study_blog/repository"
	"j_study_blog/web"

	"github.com/gofiber/fiber/v2"
)

type VocabController struct {
	vocabRepo repository.IVocabRepo
}

func NewVocabController(vocabRepo repository.IVocabRepo) VocabController {
	return VocabController{vocabRepo}
}

type FindVocabRequest struct {
	Kanji   string `form:"kanji_reading"`
	Meaning string `form:"meaning"`
	Reading string `form:"reading"`
}

func (c *VocabController) FindVocab(ctx *fiber.Ctx) error {
	request := new(FindVocabRequest)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	filter := dictionary.Vocab{KanjiReading: request.Kanji}

	if request.Meaning != "" {
		filter.Meanings = append(filter.Meanings, dictionary.VocabMeaning{Text: request.Meaning})
	}

	if request.Reading != "" {
		filter.KanaReadings = append(filter.KanaReadings, request.Reading)
	}

	vocab, err := c.vocabRepo.FindBy(filter)

	if err != nil {
		return ctx.JSON(web.NewClientError("no results matching your query for '%v'", ctx.Params("kanji")))
	} else {
		return ctx.JSON(vocab)
	}
}
