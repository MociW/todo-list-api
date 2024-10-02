package repository

import (
	"todo-list-api/internal/entity"

	"gorm.io/gorm"
)

type TodoRepository struct {
	Repository[entity.Todo]
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (t *TodoRepository) FindTodo(db *gorm.DB, todo *entity.Todo, userId string, id any) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(todo).Error
}

func (t *TodoRepository) FindAllTodo(db *gorm.DB, userId string) ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
