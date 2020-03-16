package api_v1

import (
	"encoding/json"
	"io/ioutil"
	"j_study_blog/dictionary"
	"j_study_blog/tests"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
)

func TestGetBy(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	assertions := map[string]struct {
		contentType string
		postData    url.Values
		repoExpects dictionary.Vocab
		want        []dictionary.Vocab
	}{
		"posting kanji reading": {
			contentType: "application/x-www-form-urlencoded",
			postData:    url.Values{"kanji_reading": []string{"someKanji"}},
			repoExpects: dictionary.Vocab{KanjiReading: "someKanji"},
			want:        []dictionary.Vocab{dictionary.NewVocab(), dictionary.NewVocab()},
		},
		"posting meanings": {
			contentType: "application/x-www-form-urlencoded",
			postData:    url.Values{"meanings": []string{"ice", "frost"}},
			repoExpects: dictionary.Vocab{Meanings: []dictionary.VocabMeaning{
				dictionary.VocabMeaning{Text: "ice"},
				dictionary.VocabMeaning{Text: "frost"},
			}},
			want:        []dictionary.Vocab{dictionary.Vocab{KanjiReading: "kanji for frost"}, dictionary.Vocab{KanjiReading: "kanji for ice"}},
		},
	}

	for name, tc := range assertions {
		t.Run(name, func(t *testing.T) {
			mockRepo := new(tests.MockVocabRepo)
			mockRepo.On("FindBy", tc.repoExpects).Return(tc.want, nil)
			controller := NewVocabController(mockRepo)
			app.Post("/vocabs", controller.FindVocab)
			req, _ := http.NewRequest("POST", "/vocabs", strings.NewReader(tc.postData.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Content-Length", strconv.Itoa(len(tc.postData.Encode())))

			resp, _ := app.Test(req)
			var got []dictionary.Vocab
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &got)

			assert.Equal(got, tc.want, "got %v\nwant%v", got, tc.want)
		})
	}
}
