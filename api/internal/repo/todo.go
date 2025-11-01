package repo

import (
	"github.com/ktsu2i/connect-todo/api/internal/model"
	"gorm.io/gorm"
)

type ITodoRepo interface {
	Create(title string) (*model.Todo, error)
}

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) Create(title string) (*model.Todo, error) {
	todo := &model.Todo{
		Title: title,
		Done:  false,
	}
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}, nil
}
