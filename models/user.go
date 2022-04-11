package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName     string `json:"full_name"`
	Username     string `json:"username"`
	Password     string
	NPM          string `json:"npm"`
	LastIssuedAt int64
}
