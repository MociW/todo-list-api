package helper

import (
	"todo-list-api/internal/entity"
	"todo-list-api/internal/model"
)

func ConvertUserResponse(user *entity.User) *model.UserResponse {
	response := &model.UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return response
}
