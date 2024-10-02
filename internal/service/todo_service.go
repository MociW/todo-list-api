package service

import (
	"context"
	"todo-list-api/internal/entity"
	"todo-list-api/internal/helper"
	"todo-list-api/internal/model"
	"todo-list-api/internal/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoService struct {
	DB             *gorm.DB
	TodoRepository *repository.TodoRepository
	UserRepository *repository.UserRepository
}

func NewTodoService(db *gorm.DB, todoRepository *repository.TodoRepository, userRepository *repository.UserRepository) *TodoService {
	return &TodoService{
		DB:             db,
		TodoRepository: todoRepository,
		UserRepository: userRepository,
	}
}

func (s *TodoService) Create(ctx context.Context, request *model.CreateTodoRequest) (*model.TodoResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := s.UserRepository.FindByUUID(tx, user, request.UserId); err != nil {
		return nil, fiber.ErrNotFound
	}

	todo := &entity.Todo{
		UserId: request.UserId,
		Todo:   request.Todo,
	}

	if err := s.TodoRepository.Create(tx, todo); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertTodoResponse(todo), nil
}

func (s *TodoService) Update(ctx context.Context, request *model.UpdateTodoRequest) (*model.TodoResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo := new(entity.Todo)
	if err := s.TodoRepository.FindTodo(tx, todo, request.UserId, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	todo.Todo = request.Todo

	if err := s.TodoRepository.Update(tx, todo); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertTodoResponse(todo), nil
}

func (s *TodoService) UpdateStatus(ctx context.Context, request *model.UpdateStatusTodoRequest) (*model.TodoResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo := new(entity.Todo)
	if err := s.TodoRepository.FindTodo(tx, todo, request.UserId, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	todo.Status = request.Status

	if err := s.TodoRepository.Update(tx, todo); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertTodoResponse(todo), nil
}

func (s *TodoService) Delete(ctx context.Context, request *model.DeleteTodoRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo := new(entity.Todo)
	if err := s.TodoRepository.FindTodo(tx, todo, request.UserId, request.ID); err != nil {
		return fiber.ErrBadRequest
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	err := s.TodoRepository.Delete(tx, todo)

	return err
}

func (s *TodoService) Get(ctx context.Context, request *model.GetTodoRequest) (*model.TodoResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo := new(entity.Todo)
	err := s.TodoRepository.FindTodo(tx, todo, request.UserId, request.ID)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertTodoResponse(todo), nil
}

func (s *TodoService) List(ctx context.Context, request *model.ListTodoRequest) ([]model.TodoResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todos, err := s.TodoRepository.FindAllTodo(tx, request.UserId)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = *helper.ConvertTodoResponse(&todo)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return responses, nil
}
