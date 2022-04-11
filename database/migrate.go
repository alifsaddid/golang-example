package database

import (
	"Oauth/models"
	"fmt"
)

func Migrate() {
	toMigrate := []interface{}{
		&models.Client{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
	}

	db, err := GetPostgresClient()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	db.AutoMigrate(toMigrate...)
}
