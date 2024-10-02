package entity

import "time"

type Todo struct {
	ID        uint      `gorm:"column:id"`
	UserId    string    `gorm:"column:user_id"`
	Todo      string    `gorm:"column:todo"`
	Status    int8      `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
}

func (t *Todo) TableName() string {
	return "todos"
}
