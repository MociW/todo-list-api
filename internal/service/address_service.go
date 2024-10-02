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

type AddressService struct {
	DB                *gorm.DB
	AddressRepository *repository.AddressRepository
	UserRepository    *repository.UserRepository
}

func NewAddressService(db *gorm.DB, addressRepository *repository.AddressRepository, userRepository *repository.UserRepository) *AddressService {
	return &AddressService{db, addressRepository, userRepository}
}

func (s *AddressService) Get(ctx context.Context, request *model.GetAddressRequest) (*model.AddressResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address := new(entity.Address)
	err := s.AddressRepository.FindAddress(tx, address, request.UserId, request.ID)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	responses := helper.ConvertAddressResponse(address)

	return responses, nil
}

func (s *AddressService) List(ctx context.Context, request *model.ListAddressRequest) ([]model.AddressResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	addresses, err := s.AddressRepository.FindAllAddress(tx, request.UserId)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.AddressResponse, len(addresses))
	for i, address := range addresses {
		responses[i] = *helper.ConvertAddressResponse(&address)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return responses, nil
}

func (s *AddressService) Create(ctx context.Context, request *model.CreateAddressRequest) (*model.AddressResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := s.UserRepository.FindByUUID(tx, user, request.UserId); err != nil {
		return nil, fiber.ErrNotFound
	}

	address := &entity.Address{
		UserId:     user.UUID,
		Street:     request.Street,
		City:       request.City,
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	if err := s.AddressRepository.Create(tx, address); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertAddressResponse(address), nil
}

func (s *AddressService) Update(ctx context.Context, request *model.UpdateAddressRequest) (*model.AddressResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address := new(entity.Address)
	if err := s.AddressRepository.FindAddress(tx, address, request.UserId, request.ID); err != nil {
		return nil, fiber.ErrNotFound
	}

	address.Street = request.Street
	address.City = request.City
	address.Country = request.Country
	address.PostalCode = request.PostalCode

	if err := s.AddressRepository.Update(tx, address); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return helper.ConvertAddressResponse(address), nil
}

func (s *AddressService) Delete(ctx context.Context, request *model.DeleteAddressRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address := new(entity.Address)
	if err := s.AddressRepository.FindAddress(tx, address, request.UserId, request.ID); err != nil {
		return fiber.ErrBadRequest
	}

	if err := tx.Commit().Error; err != nil {
		return fiber.ErrInternalServerError
	}

	err := s.AddressRepository.Delete(tx, address)

	return err
}
