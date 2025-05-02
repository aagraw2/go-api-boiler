package repository

import (
	"go-api-boiler/models"
	"go-api-boiler/utils"
	"sync"
)

type UserRepository interface {
	GetUsers() []models.User
	CreateUser(user models.User) models.User
}

type userRepo struct {
	db *utils.InMemoryDB
	mu sync.Mutex
}

func NewUserRepository(db *utils.InMemoryDB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUsers() []models.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.db.Users
}

func (r *userRepo) CreateUser(user models.User) models.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = uint64(len(r.db.Users) + 1) // Auto-increment
	r.db.Users = append(r.db.Users, user)
	return user
}
