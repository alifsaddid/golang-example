package services

import (
	"Oauth/models"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthService interface {
	Login(username string, password string, clientId string, clientSecret string) (models.Token, models.Token, *models.Error)
	Resource(token string) (models.User, models.Token, *models.Error)
}

type authService struct {
	jwtService    JWTService
	userService   UserService
	redisService  RedisService
	clientService ClientService
}

func NewAuthService(jwtService JWTService, userService UserService, redisService RedisService, clientService ClientService) AuthService {
	return &authService{
		jwtService:    jwtService,
		userService:   userService,
		redisService:  redisService,
		clientService: clientService,
	}
}

func (a authService) Login(username string, password string, clientId string, clientSecret string) (models.Token, models.Token, *models.Error) {
	_, err := a.clientService.GetClient(clientId, clientSecret)
	if err != nil {
		return models.Token{}, models.Token{}, &models.Error{
			Err:    errors.New("client not found"),
			Status: http.StatusNotFound,
		}
	}

	user, err := a.userService.AuthenticateUser(username, password)
	if err != nil {
		return models.Token{}, models.Token{}, &models.Error{
			Err:    errors.New("user not authenticated"),
			Status: http.StatusUnauthorized,
		}
	}

	// Generate Token
	accessToken := a.jwtService.Generate(int(user.ID), clientId, clientSecret, "access")
	refreshToken := a.jwtService.Generate(int(user.ID), clientId, clientSecret, "refresh")
	accessToken.RefreshToken = refreshToken.Token

	t, _ := json.Marshal(accessToken)
	stringToken := string(t)
	accessToken.TokenJson = stringToken

	t, _ = json.Marshal(refreshToken)
	stringToken = string(t)
	refreshToken.TokenJson = stringToken

	// Save to redis
	a.redisService.WriteCache(accessToken.Token, accessToken.TokenJson)
	a.redisService.WriteCache(refreshToken.Token, refreshToken.TokenJson)

	return accessToken, refreshToken, nil
}

func (a authService) Resource(token string) (models.User, models.Token, *models.Error) {
	data, t, err := a.jwtService.Verify(token)
	if err != nil {
		return models.User{}, models.Token{}, &models.Error{
			Err:    err.Err,
			Status: err.Status,
		}
	}

	return data, t, nil
}
