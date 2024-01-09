package db

import (
	"fmt"
	"todos/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(env *config.Env) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		env.DB_HOST, env.DB_USER, env.DB_PASS, env.DB_NAME, env.DB_PORT)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func RunAutoMigartion() {
	if err := DB.AutoMigrate(&Todo{}); err != nil {
		panic("Failed to run migration!")
	}
}
