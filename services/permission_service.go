package services

import (
	"Oauth/models"
	"Oauth/repositories"
)

type PermissionService interface {
	GetAll() ([]*models.Permission, error)
	GetByIds(ids []int) ([]*models.Permission, error)
}

type permissionService struct {
	permissionRepository repositories.PermissionRepository
}

func NewPermissionService(p repositories.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: p,
	}
}

func (p permissionService) GetAll() ([]*models.Permission, error) {
	return p.permissionRepository.GetAll()
}

func (p permissionService) GetByIds(ids []int) ([]*models.Permission, error) {
	return p.permissionRepository.GetAllByIds(ids)
}
