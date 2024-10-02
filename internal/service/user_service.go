package service

import (
	"context"
	"todo-list-api/internal/entity"
	"todo-list-api/internal/helper"
	"todo-list-api/internal/model"
	"todo-list-api/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB             *gorm.DB
	UserRepository *repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository *repository.UserRepository) *UserService {
	return &UserService{DB: db, UserRepository: userRepository}
}

func (s *UserService) Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, err := s.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	if total > 0 {
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	token := uuid.New().String()

	user := &entity.User{
		Username: request.Username,
		UUID:     token,
		Email:    request.Email,
		Password: string(password),
	}

	if err := s.UserRepository.Create(tx, user); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertUserResponse(user), nil
}

func (s *UserService) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := s.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	return helper.ConvertUserResponse(user), nil
}

func (s *UserService) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := s.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		return nil, fiber.ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = string(password)

	if err := s.UserRepository.Update(tx, user); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertUserResponse(user), nil
}

func (s *UserService) Delete(ctx context.Context, request *model.DeleteUserRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := s.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		return fiber.ErrBadRequest
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	err := s.UserRepository.Delete(tx, user)

	return err
}
