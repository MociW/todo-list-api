package model

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
