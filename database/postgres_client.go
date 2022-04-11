package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func GetPostgresConfig() Config {
	return Config{
		DBUser: os.Getenv("DB_USERNAME"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
}

func GetPostgresClient() (*gorm.DB, error) {
	cfg := GetPostgresConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
