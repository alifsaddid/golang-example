package services

import (
	"Oauth/dto/request"
	"Oauth/models"
	"Oauth/repositories"
	"errors"
	"fmt"
)

type RoleService interface {
	Create(role request.RoleRequest) (models.Role, error)
	UpdateById(id int, role request.RoleRequest) (models.Role, error)
	DeleteById(id int) error
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
}

type roleService struct {
	roleRepository    repositories.RoleRepository
	permissionService PermissionService
}

func NewRoleService(r repositories.RoleRepository, p PermissionService) RoleService {
	return &roleService{
		roleRepository:    r,
		permissionService: p,
	}
}

func (r roleService) Create(role request.RoleRequest) (models.Role, error) {
	_, err := r.roleRepository.GetByName(role.Name)
	if err == nil {
		return models.Role{}, fmt.Errorf("role with name %s already exists", role.Name)
	}

	permissions, err := r.permissionService.GetByIds(role.Permissions)
	if err != nil {
		return models.Role{}, err
	}
	if len(permissions) != len(role.Permissions) {
		return models.Role{}, errors.New("role ids invalid")
	}

	roleModel := models.Role{
		Name:        role.Name,
		Permissions: permissions,
	}
	return r.roleRepository.Create(roleModel)
}

func (r roleService) UpdateById(id int, role request.RoleRequest) (models.Role, error) {
	currentRole, err := r.roleRepository.GetByName(role.Name)
	if err == nil && currentRole.ID != uint(id) {
		return models.Role{}, fmt.Errorf("role with name %s already exists", role.Name)
	}

	permissions, err := r.permissionService.GetByIds(role.Permissions)
	if err != nil {
		return models.Role{}, err
	}
	if len(permissions) != len(role.Permissions) {
		return models.Role{}, errors.New("role ids invalid")
	}

	roleModel, err := r.roleRepository.GetById(id)
	if err != nil {
		return models.Role{}, err
	}

	roleModel.Name = role.Name
	roleModel.Permissions = permissions
	return r.roleRepository.Update(roleModel)
}

func (r roleService) DeleteById(id int) error {
	return r.roleRepository.Delete(id)
}

func (r roleService) GetAll() ([]models.Role, error) {
	return r.roleRepository.GetAll()
}

func (r roleService) GetById(id int) (models.Role, error) {
	return r.roleRepository.GetById(id)
}
