package model

import "time"

type AddressResponse struct {
	ID         uint      `json:"id"`
	UserId     string    `json:"user_id"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	PostalCode string    `json:"postal_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type ListAddressRequest struct {
	UserId string `json:"user_id"`
}

type GetAddressRequest struct {
	ID     uint   `json:"-"`
	UserId string `json:"user_id"`
}

type CreateAddressRequest struct {
	UserId     string `json:"user_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type UpdateAddressRequest struct {
	ID         uint   `json:"id"`
	UserId     string `json:"user_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type DeleteAddressRequest struct {
	ID     uint   `json:"-"`
	UserId string `json:"user_id"`
}
