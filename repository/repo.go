package repository

import (
	"errors"
	"j_study_blog/dictionary"
)

type InsertGuid = string

type RepoError error

var (
	InsertError = errors.New("unable to insert model into repo")
)

//IVocabRepo interface for querying for dictionary.Vocab
type IVocabRepo interface {
	Insert(dictionary.Vocab) (InsertGuid, RepoError)
	InsertMany([]dictionary.Vocab) RepoError
	FindBy(filter dictionary.Vocab) ([]dictionary.Vocab, RepoError)
	List(offset int64, limit int64) ([]dictionary.Vocab, RepoError)
}
