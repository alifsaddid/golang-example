package repositories

import (
	"Oauth/models"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetAllByIds(ids []int) ([]*models.Permission, error)
	GetAll() ([]*models.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (p permissionRepository) GetAllByIds(ids []int) ([]*models.Permission, error) {
	var perms []*models.Permission
	res := p.db.Where("id IN ?", ids).Find(&perms)
	return perms, res.Error
}

func (p permissionRepository) GetAll() ([]*models.Permission, error) {
	var perms []*models.Permission
	res := p.db.Find(&perms)
	return perms, res.Error
}
