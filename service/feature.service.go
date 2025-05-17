package service

import (
	"fmt"
	"go-api-boiler/models"
	"go-api-boiler/repository"
)

type FeatureService interface {
	UpdateFeature(sub models.Subscription, feature models.Feature, action models.FeatureState) (bool, error)
}

type FeatureServ struct {
	repo       repository.FeatureRepository
	clientRepo repository.ClientRepository
	depRepo    repository.DependencyRepository
}

func NewFeatureService(FeatureRepo repository.FeatureRepository, clientRepo repository.ClientRepository, depRepo repository.DependencyRepository) FeatureService {
	return &FeatureServ{repo: FeatureRepo, clientRepo: clientRepo, depRepo: depRepo}
}

func (s *FeatureServ) UpdateFeature(sub models.Subscription, feature models.Feature, action models.FeatureState) (bool, error) {
	if action == models.ENABLE {
		enabledFeatureMap := sub.GetEnabledFeatures()
		dependencies := s.depRepo.GetDependencies(feature)
		for c := range dependencies {
			_, found := enabledFeatureMap[c]
			if !found {
				return false, fmt.Errorf("one or more dependencies not enabled")
			}
		}
		return s.repo.EnableFeature(sub, feature)
	} else if action == models.DISABLE {
		dependents := s.depRepo.GetDependents(feature)

		for c := range dependents {
			s.repo.DisableFeature(sub, c)
		}
		return true, nil
	}
	return true, nil
}
