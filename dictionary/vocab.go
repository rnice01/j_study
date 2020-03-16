package dictionary

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vocab struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	KanjiReading string `bson:"kanji_reading" json:"kanji_reading"`
	Meanings []VocabMeaning `bson:"meanings" json:"meanings"`
	KanaReadings []string `bson:"kana_readings" json:"kana_readings"`
}

type VocabMeaning struct {
	Text string `bson:"text" json:"text"`
	Language string `bson:"language" json:"language"`
}

func NewVocab() Vocab {
	return Vocab{
		KanjiReading: "",
		Meanings:     []VocabMeaning{},
		KanaReadings: []string{},
	}
}
