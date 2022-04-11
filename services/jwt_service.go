package services

import (
	"Oauth/models"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type JWTService interface {
	Generate(uid int, clientId string, clientSecret string, tp string) models.Token
	Verify(token string) (models.User, models.Token, *models.Error)
}

type jwtService struct {
	redisService  RedisService
	userService   UserService
	clientService ClientService
}

func NewJwtService(redisService RedisService, userService UserService, clientService ClientService) JWTService {
	return &jwtService{
		redisService:  redisService,
		userService:   userService,
		clientService: clientService,
	}
}

func (j jwtService) Generate(uid int, clientId string, clientSecret string, tp string) models.Token {
	issuedTime := time.Now().Unix()
	token := models.Token{
		UserId:       uid,
		ClientId:     clientId,
		ClientSecret: clientId,
		IssuedAt:     issuedTime,
		Type:         tp,
	}
	t, _ := json.Marshal(token)
	stringToken := string(t)
	sha := sha1.New()
	sha.Write([]byte(stringToken))
	hashed := fmt.Sprintf("%x", string(sha.Sum(nil)))
	token.Token = string(hashed)

	return token
}

func (j jwtService) Verify(token string) (models.User, models.Token, *models.Error) {
	t, found := j.redisService.GetCache(token)

	if !found {
		return models.User{}, models.Token{}, &models.Error{
			Err:    errors.New("token invalid"),
			Status: http.StatusUnauthorized,
		}
	}

	var data models.Token
	json.Unmarshal([]byte(t), &data)

	user, err := j.userService.GetUserById(data.UserId)
	if err != nil {
		return models.User{}, models.Token{}, &models.Error{
			Err:    errors.New("user not found"),
			Status: http.StatusNotFound,
		}
	}

	issuedTime := time.Unix(data.IssuedAt, 0)
	currentTime := time.Now()
	if currentTime.After(issuedTime.Add(time.Minute * 5)) {
		return models.User{}, models.Token{}, &models.Error{
			Err:    errors.New("token expired"),
			Status: http.StatusUnauthorized,
		}
	}

	return user, data, nil
}
