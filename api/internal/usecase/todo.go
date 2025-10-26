package usecase

import (
	"errors"

	"github.com/ktsu2i/connect-todo/api/internal/model"
	"github.com/ktsu2i/connect-todo/api/internal/repo"
)

var ErrEmptyTitle = errors.New("title must not be empty")

type ITodoUsecase interface {
	Create(title string) (*model.Todo, error)
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
