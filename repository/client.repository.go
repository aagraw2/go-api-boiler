package repository

import (
	"go-api-boiler/models"
)

type ClientRepository struct {
	Clients []models.Client
}

func NewClientRepository() ClientRepository {
	return ClientRepository{
		Clients: []models.Client{},
	}
}
