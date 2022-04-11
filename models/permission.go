package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string `json:"name"`
	Code  string `json:"code"`
	Group string `json:"group"`
}
