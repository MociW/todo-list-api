package repository

import (
	"todo-list-api/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email any) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *Repository[T]) CountByEmail(db *gorm.DB, email any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("email = ?", email).Count(&total).Error
	return total, err
}
