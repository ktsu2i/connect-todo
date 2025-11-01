package repo

import (
	"errors"

	"github.com/ktsu2i/connect-todo/api/internal/model"
	"gorm.io/gorm"
)

type ITodoRepo interface {
	Create(title string) (*model.Todo, error)
	Update(id int64, title string, done bool) (*model.Todo, error)
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

func (r *TodoRepo) Update(id int64, title string, done bool) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.Where("id = ?", id).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if title != "" {
		todo.Title = title
	}
	todo.Done = done

	if err := r.db.Save(&todo).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
