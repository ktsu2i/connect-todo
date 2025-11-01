package usecase

import (
	"errors"

	"github.com/ktsu2i/connect-todo/api/internal/model"
	"github.com/ktsu2i/connect-todo/api/internal/repo"
)

var ErrEmptyTitle = errors.New("title must not be empty")

type ITodoUsecase interface {
	Create(title string) (*model.Todo, error)
	Update(id int64, title string, done bool) (*model.Todo, error)
}

type todoUsecase struct {
	repo repo.ITodoRepo
}

func NewTodoUsecase(repo repo.ITodoRepo) ITodoUsecase {
	return &todoUsecase{repo: repo}
}

func (uc *todoUsecase) Create(title string) (*model.Todo, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return uc.repo.Create(title)
}

func (uc *todoUsecase) Update(id int64, title string, done bool) (*model.Todo, error) {
	todo, err := uc.repo.Update(id, title, done)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil
	}
	return todo, nil
}
