package service

import (
	"go-api-boiler/models"
	"go-api-boiler/repository"
)

type DependencyService interface {
	UpdateDependency(parent, child models.Feature) (bool, error)
}

type DependencyServ struct {
	repo repository.DependencyRepository
}

func NewDependencyService(DependencyRepo repository.DependencyRepository) DependencyService {
	return &DependencyServ{repo: DependencyRepo}
}

func (s *DependencyServ) UpdateDependency(parent, child models.Feature) (bool, error) {
	return s.repo.AddDependency(parent, child)
}
