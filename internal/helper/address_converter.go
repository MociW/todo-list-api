package helper

import (
	"todo-list-api/internal/entity"
	"todo-list-api/internal/model"
)

func ConvertAddressResponse(address *entity.Address) *model.AddressResponse {
	response := &model.AddressResponse{
		ID:         address.ID,
		UserId:     address.UserId,
		Street:     address.Street,
		City:       address.City,
		Country:    address.Country,
		PostalCode: address.PostalCode,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}
	return response
}
