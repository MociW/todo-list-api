package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID      string    `gorm:"column:uuid"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Todos     []Todo    `gorm:"foreignKey:user_id;references:id"`
	Addresses []Address `gorm:"foreignKey:user_id;references:id"`
}
