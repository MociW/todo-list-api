package model

import "time"

type TodoResponse struct {
	ID        uint      `json:"id"`
	UserId    string    `json:"user_id"`
	Todo      string    `json:"todo"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	UserId string `json:"user_id"`
	Todo   string `json:"todo"`
}

type UpdateTodoRequest struct {
	ID     uint   `json:"id"`
	UserId string `json:"user_id"`
	Todo   string `json:"todo"`
}

type DeleteTodoRequest struct {
	ID     uint   `json:"id"`
	UserId string `json:"user_id"`
}

type GetTodoRequest struct {
	ID     uint   `json:"id"`
	UserId string `json:"user_id"`
}

type ListTodoRequest struct {
	UserId string `json:"user_id"`
}

type UpdateStatusTodoRequest struct {
	ID     uint   `json:"id"`
	UserId string `json:"user_id"`
	Status int8   `json:"status"`
}
