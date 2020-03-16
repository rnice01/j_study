package repository

import (
	"j_study_blog/dictionary"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestVocabToBsonFilter(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		vocab dictionary.Vocab
		want  bson.M
	}{
		"meanings only": {
			vocab: dictionary.Vocab{
				Meanings: []dictionary.VocabMeaning{
					dictionary.VocabMeaning{Text: "first"},
					dictionary.VocabMeaning{Text: "second"},
				},
			},
			want: bson.M{"meanings": bson.M{"$in": bson.A{"first", "second"}}},
		},
		"kanji reading": {
			vocab: dictionary.Vocab{
				KanjiReading: "some kanji",
			},
			want: bson.M{"kanji_reading": "some kanji"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := VocabToBsonFilter(tc.vocab)
			assert.Equal(got, tc.want, "got %v\nwant%v", got, tc.want)
		})
	}
}
