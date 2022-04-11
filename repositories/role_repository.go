package repositories

import (
	"Oauth/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(r models.Role) (models.Role, error)
	Update(r models.Role) (models.Role, error)
	Delete(id int) error
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
	GetByName(name string) (models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(role models.Role) (models.Role, error) {
	res := r.db.Create(&role)
	return role, res.Error
}

func (r *roleRepository) Update(role models.Role) (models.Role, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Save(&role)
		if res.Error != nil {
			return res.Error
		}
		err := tx.Model(&role).Association("Permissions").Replace(&role.Permissions)
		return err
	})

	return role, err
}

func (r *roleRepository) Delete(id int) error {
	res := r.db.Delete(&models.Role{}, id)
	return res.Error
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	res := r.db.Preload("Permissions").Find(&roles)
	return roles, res.Error
}

func (r *roleRepository) GetById(id int) (models.Role, error) {
	var role models.Role
	res := r.db.Preload("Permissions").Where("id = ?", id).First(&role)
	return role, res.Error
}

func (r *roleRepository) GetByName(name string) (models.Role, error) {
	var role models.Role
	res := r.db.Where("name = ?", name).First(&role)
	return role, res.Error
}
