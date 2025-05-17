package repository

import (
	"go-api-boiler/models"
	"sync"
)

type SubscriptionRepository interface {
	EnableFeature(sub models.Subscription, feature models.Feature) (bool, error)
	DisableFeature(sub models.Subscription, feature models.Feature) (bool, error)
}

type subRepo struct {
	subs []models.Subscription
	mu   sync.Mutex
}

func NewSubscriptionRepository(c ClientRepository) SubscriptionRepository {
	return &subRepo{
		subs: []models.Subscription{},
	}
}

func (r *subRepo) EnableFeature(sub models.Subscription, feature models.Feature) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sub.EnableFeature(feature)
	return false, nil
}

func (r *subRepo) DisableFeature(sub models.Subscription, feature models.Feature) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sub.DisableFeature(feature)
	return false, nil
}
