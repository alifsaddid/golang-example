package services

import (
	"Oauth/models"
	"Oauth/repositories"
)

type ClientService interface {
	GetClient(clientId string, clientSecret string) (models.Client, error)
}

type clientService struct {
	clientRepository repositories.ClientRepository
}

func NewClientService(clientRepository repositories.ClientRepository) ClientService {
	return &clientService{
		clientRepository: clientRepository,
	}
}

func (c clientService) GetClient(clientId string, clientSecret string) (models.Client, error) {
	return c.clientRepository.FindByClientIdAndClientSecret(clientId, clientSecret)
}
