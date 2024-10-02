package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/golang_todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}

// migrate -database "mysql://root:@tcp(localhost:3306)/golang_todo_list" -path db/migrations up
// migrate -database "mysql://root:@tcp(localhost:3306)/golang_todo_list" -path db/migrations down
// migrate create -ext sql -dir db/migrations create_table_todos
// migrate create -ext sql -dir db/migrations create_table_addresses
// migrate create -ext sql -dir db/migrations create_table_users
