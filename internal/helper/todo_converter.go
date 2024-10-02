package helper

import (
	"todo-list-api/internal/entity"
	"todo-list-api/internal/model"
)

func ConvertTodoResponse(todo *entity.Todo) *model.TodoResponse {
	response := &model.TodoResponse{
		ID:        todo.ID,
		UserId:    todo.UserId,
		Todo:      todo.Todo,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return response
}
