package repositories

import (
	"Oauth/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id int) (models.User, error)
	FindByUsernameAndPassword(username string, password string) (models.User, error)
	Update(models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) FindById(id int) (models.User, error) {
	var user models.User
	err := u.db.Model(&models.User{}).Where("id = ?", id).First(&user)
	return user, err.Error
}

func (u userRepository) FindByUsernameAndPassword(username string, password string) (models.User, error) {
	var user models.User
	err := u.db.Model(&models.User{}).Where("username = ? AND password = ?", username, password).First(&user)
	return user, err.Error
}

func (u userRepository) Update(user models.User) (models.User, error) {
	res := u.db.Save(&user)
	return user, res.Error
}
