package database

import (
	"app/config"
	"app/models"
	"fmt"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	var err error

	dbName := config.Config("DB_NAME")

	DB, err = gorm.Open(sqlite.Open(dbName+".db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	migrateAll()
}

func migrateAll() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}

	fmt.Println("Database migrated successfully")
}
