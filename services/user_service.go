package services

import (
	"Oauth/models"
	"Oauth/repositories"
	"crypto/sha1"
	"fmt"
)

type UserService interface {
	GetUserById(id int) (models.User, error)
	AuthenticateUser(username string, password string) (models.User, error)
	UpdateLastIssuedTime(id int, lastIssued int64) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) GetUserById(id int) (models.User, error) {
	return u.userRepository.FindById(id)
}

func (u userService) AuthenticateUser(username string, password string) (models.User, error) {
	sha := sha1.New()
	sha.Write([]byte(password))
	hashed := fmt.Sprintf("%x", string(sha.Sum(nil)))
	return u.userRepository.FindByUsernameAndPassword(username, hashed)
}

func (u userService) UpdateLastIssuedTime(id int, lastIssued int64) (models.User, error) {
	user, _ := u.GetUserById(id)
	user.LastIssuedAt = lastIssued
	return u.userRepository.Update(user)
}
