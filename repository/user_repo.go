package repository

import (
	"go-api-boiler/models"
	"sync"
)

type UserRepository interface {
	GetUsers() []models.User
	CreateUser(user models.User) models.User
}

type userRepo struct {
	users []models.User
	mu    sync.Mutex
}

func NewUserRepository() UserRepository {
	return &userRepo{
		users: []models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
		},
	}
}

func (r *userRepo) GetUsers() []models.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.users
}

func (r *userRepo) CreateUser(user models.User) models.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = uint64(len(r.users) + 1) // auto-increment
	r.users = append(r.users, user)
	return user
}
