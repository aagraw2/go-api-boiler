package utils

import (
	"go-api-boiler/models"
	"sync"
)

type InMemoryDB struct {
	Users []models.User
}

var (
	instance *InMemoryDB
	once     sync.Once
)

// GetDBInstance returns the singleton instance of InMemoryDB
func GetDBInstance() *InMemoryDB {
	once.Do(func() {
		instance = &InMemoryDB{
			Users: make([]models.User, 0),
		}
	})
	return instance
}
