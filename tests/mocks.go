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

func (r *MockVocabRepo) List(offset int64, limit int64) ([]dictionary.Vocab, repository.RepoError) {
	args := r.Called(offset, limit)
	arg1 := args.Get(0).([]dictionary.Vocab)
	arg2 := args.Get(1)

	if arg2 == nil {
		return arg1, nil
	}
	return arg1, arg2.(repository.RepoError)
}

func (r *MockVocabRepo) FindMany(filter dictionary.Vocab) ([]dictionary.Vocab, repository.RepoError) {
	return nil, nil
}
