package repository

import (
	"go-api-boiler/models"
	"sync"
)

type FeatureRepository interface {
	EnableFeature(sub models.Subscription, feature models.Feature) (bool, error)
	DisableFeature(sub models.Subscription, feature models.Feature) (bool, error)
}

type featureRepo struct {
	features []models.Feature
	mu       sync.Mutex
}

func NewFeatureRepository(c ClientRepository) FeatureRepository {
	return &featureRepo{
		features: []models.Feature{},
	}
}

func (r *featureRepo) EnableFeature(sub models.Subscription, feature models.Feature) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sub.EnableFeature(feature)
	return false, nil
}

func (r *featureRepo) DisableFeature(sub models.Subscription, feature models.Feature) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sub.DisableFeature(feature)
	return false, nil
}
