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

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestGetBy(t *testing.T) {
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
			postData:    url.Values{"meaning": []string{"frost"}},
			repoExpects: dictionary.Vocab{Meanings: []dictionary.VocabMeaning{dictionary.VocabMeaning{Text: "frost"}}},
			want:        []dictionary.Vocab{dictionary.Vocab{KanjiReading: "kanji for frost"}, dictionary.Vocab{KanjiReading: "kanji for ice"}},
		},
		"posting readings": {
			contentType: "application/x-www-form-urlencoded",
			postData:    url.Values{"reading": []string{"kanaReading"}},
			repoExpects: dictionary.Vocab{KanaReadings: []string{"kanaReading"}},
			want:        []dictionary.Vocab{dictionary.Vocab{KanaReadings: []string{"kanaReading"}}},
		},
	}

	for name, tc := range assertions {
		t.Run(name, func(t *testing.T) {
			app := fiber.New()
			assert := assert.New(t)
			mockRepo := new(tests.MockVocabRepo)
			mockRepo.On("FindBy", tc.repoExpects).Return(tc.want, nil)
			controller := NewVocabController(mockRepo)
			app.Post("/vocabs", controller.FindVocab)
			req, _ := http.NewRequest("POST", "/vocabs", strings.NewReader(tc.postData.Encode()))
			req.Header.Add("Content-Type", tc.contentType)
			req.Header.Add("Content-Length", strconv.Itoa(len(tc.postData.Encode())))

			resp, _ := app.Test(req)
			var got []dictionary.Vocab
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &got)

			assert.Equal(got, tc.want, "got %v\nwant%v", got, tc.want)
		})
	}
}

func TestListVocab(t *testing.T) {
	assertions := map[string]struct {
		queryString       string
		repoExpectsOffset int64
		repoExpectsLimit  int64
	}{
		"no query string": {
			queryString:       "",
			repoExpectsOffset: 0,
			repoExpectsLimit:  25,
		},
		"limit in query string": {
			queryString: "?limit=250",
			repoExpectsLimit: 250,
			repoExpectsOffset: 0,
		},
		"offset in query string": {
			queryString: "?offset=500",
			repoExpectsOffset: 500,
			repoExpectsLimit: 25,
		},
		"limit and offset in query string": {
			queryString: "?limit=10&offset=50",
			repoExpectsLimit: 10,
			repoExpectsOffset: 50,
		},
	}

	for name, tc := range assertions {
		t.Run(name, func(t *testing.T) {
			app := fiber.New()
			mockRepo := new(tests.MockVocabRepo)
			var mock []dictionary.Vocab
			mockRepo.On("List", tc.repoExpectsOffset, tc.repoExpectsLimit).Return(mock, nil)
			controller := NewVocabController(mockRepo)
			app.Get("/vocabs", controller.ListVocab)
			req, _ := http.NewRequest("GET", "/vocabs"+tc.queryString, strings.NewReader(""))

			resp, _ := app.Test(req)
			var got []dictionary.Vocab
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &got)

			mockRepo.AssertExpectations(t)
		})
	}
}
