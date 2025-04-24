package service

import (
	"go-api-boiler/models"
	"go-api-boiler/repository"
)

type UserService interface {
	GetUsers() []models.User
	CreateUser(user models.User) models.User
}

type userServ struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServ{repo: userRepo}
}

func (s *userServ) GetUsers() []models.User {
	return s.repo.GetUsers()
}

func (s *userServ) CreateUser(user models.User) models.User {
	return s.repo.CreateUser(user)
}
