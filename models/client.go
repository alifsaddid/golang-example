package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientId     string
	ClientSecret string
}
