package repositories

import (
	"Oauth/models"

	"gorm.io/gorm"
)

type ClientRepository interface {
	FindByClientIdAndClientSecret(clientId string, clientSecret string) (models.Client, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{
		db: db,
	}
}

func (c clientRepository) FindByClientIdAndClientSecret(clientId string, clientSecret string) (models.Client, error) {
	var client models.Client
	err := c.db.Model(&models.Client{}).Where("client_id = ? AND client_secret = ?", clientId, clientSecret).First(&client)
	return client, err.Error
}
