package tests

import (
	"j_study_blog/dictionary"
	"j_study_blog/repository"

	"github.com/stretchr/testify/mock"
)

type MockVocabRepo struct {
	mock.Mock
}

func (r *MockVocabRepo) Insert(dictionary.Vocab) (repository.InsertGuid, repository.RepoError) {
	return "", nil
}

func (r *MockVocabRepo) InsertMany([]dictionary.Vocab) repository.RepoError {
	return nil
}

func (r *MockVocabRepo) FindBy(filter dictionary.Vocab) ([]dictionary.Vocab, repository.RepoError) {
	args := r.Called(filter)
	arg1 := args.Get(0).([]dictionary.Vocab)
	arg2 := args.Get(1)

	if arg2 == nil {
		return arg1, nil
	}
	return arg1, arg2.(repository.RepoError)
}

func (r *MockVocabRepo) FindByKanji(kanji string) (dictionary.Vocab, repository.RepoError) {
	args := r.Called(kanji)
	arg1 := args.Get(0).(dictionary.Vocab)
	arg2 := args.Get(1)

	if arg2 == nil {
		return arg1, nil
	}
	return arg1, arg2.(repository.RepoError)
}

func (r *MockVocabRepo) FindMany(filter dictionary.Vocab) ([]dictionary.Vocab, repository.RepoError) {
	return nil, nil
}
