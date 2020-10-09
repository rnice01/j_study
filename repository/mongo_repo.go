package repository

import (
	"context"
	"fmt"
	"j_study_blog/dictionary"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

type VocabRepo struct {
	db *mongo.Database
}

func initClient() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}
	db = client.Database("j_study_blog")

	return nil
}

func NewVocabRepo() *VocabRepo {
	if db == nil {
		err := initClient()
		if err != nil {
			log.Fatal(err)
		}
	}

	return &VocabRepo{db}
}

func (r *VocabRepo) Insert(vocab dictionary.Vocab) (InsertGuid, RepoError) {
	result, err := r.db.Collection("vocabs").InsertOne(context.TODO(), vocab)

	if err != nil {
		log.Fatal(err)
		return "", InsertError
	}
	return fmt.Sprintf("%s", result.InsertedID), nil
}

func (r *VocabRepo) InsertMany(vocabs []dictionary.Vocab) RepoError {

	vocabInterfaces := make([]interface{}, len(vocabs))
	for i, v := range vocabs {
		vocabInterfaces[i] = v
	}

	_, err := r.db.Collection("vocabs").InsertMany(context.TODO(), vocabInterfaces)
	if err != nil {
		log.Fatal(err)
		return InsertError
	}
	return nil
}

func (r *VocabRepo) FindBy(filter dictionary.Vocab) ([]dictionary.Vocab, RepoError) {
	var vocabs []dictionary.Vocab
	cur, err := r.db.Collection("vocabs").Find(context.TODO(), VocabToBsonFilter(filter), options.Find().SetLimit(20))

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var res dictionary.Vocab

		err := cur.Decode(&res)
		if err != nil {
			log.Fatal(err)
		}
		vocabs = append(vocabs, res)
	}

	return vocabs, nil
}

func (r *VocabRepo) FindByKanji(kanji string) (dictionary.Vocab, RepoError) {
	vocab := dictionary.NewVocab()
	res := r.db.Collection("vocabs").FindOne(context.TODO(), bson.M{"kanji_reading": kanji})

	fmt.Println(res.Err())
	dErr := res.Decode(&vocab)

	if dErr == mongo.ErrNoDocuments {
		return vocab, nil
	} else if dErr != nil {
		log.Fatal(dErr)
	}

	return vocab, nil
}

func VocabToBsonFilter(vocab dictionary.Vocab) bson.M {
	filter := bson.M{}
	if vocab.KanjiReading != "" {
		filter["kanji_reading"] = vocab.KanjiReading
	}

	if len(vocab.Meanings) > 0 {
		meanings := bson.A{}
		for _, m := range vocab.Meanings {
			meanings = append(meanings, m.Text)
		}
		filter["meanings"] = bson.M{"$in": meanings}
	}
	return filter
}
