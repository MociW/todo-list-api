package entity

import "time"

type Address struct {
	ID         uint      `gorm:"column:id"`
	UserId     string    `gorm:"column:user_id"`
	Street     string    `gorm:"column:street"`
	City       string    `gorm:"column:city"`
	Country    string    `gorm:"column:country"`
	PostalCode string    `gorm:"column:postal_code"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:user_id;references:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}
