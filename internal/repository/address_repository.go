package repository

import (
	"todo-list-api/internal/entity"

	"gorm.io/gorm"
)

type AddressRepository struct {
	Repository[entity.Address]
}

func NewAddressRepository() *AddressRepository {
	return &AddressRepository{}
}

func (a *AddressRepository) FindAddress(db *gorm.DB, address *entity.Address, userId string, id any) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(address).Error
}

func (a *AddressRepository) FindAllAddress(db *gorm.DB, userId string) ([]entity.Address, error) {
	var Addresses []entity.Address
	if err := db.Where("user_id = ?", userId).Find(&Addresses).Error; err != nil {
		return nil, err
	}
	return Addresses, nil
}
