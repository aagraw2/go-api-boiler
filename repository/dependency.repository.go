package repository

import (
	"fmt"
	"go-api-boiler/models"
	"sync"
)

type DependencyRepository interface {
	AddDependency(parent, child models.Feature) (bool, error)
	GetDependencies(root models.Feature) models.FeatureSet
	GetDependents(root models.Feature) models.FeatureSet
}

type DependencyRepo struct {
	DependencyGraph models.DependencyGraph
	mu              sync.Mutex
}

func NewDependencyRepository() DependencyRepository {
	return &DependencyRepo{
		DependencyGraph: models.DependencyGraph{
			FeaturesSet: make(models.FeatureSet),
			DependsOn:   make(models.DependencyMap),
			Dependents:  make(models.DependencyMap),
		},
	}
}

func (r *DependencyRepo) AddDependency(parent models.Feature, child models.Feature) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if child.ID == parent.ID {
		return false, fmt.Errorf("cannot depend on self")
	}
	if r.DependsOn(child, parent) {
		return false, fmt.Errorf("circular dependencies not allowed")
	}

	r.DependencyGraph.FeaturesSet[parent] = struct{}{}
	r.DependencyGraph.FeaturesSet[child] = struct{}{}

	AddRelation(r.DependencyGraph.DependsOn, child, parent)
	AddRelation(r.DependencyGraph.Dependents, parent, child)

	return true, nil
}

func AddRelation(dm models.DependencyMap, key models.Feature, node models.Feature) {
	keyNodes, found := dm[key]
	if !found {
		keyNodes = make(models.FeatureSet)
		dm[key] = keyNodes
	}
	keyNodes[node] = struct{}{}
}

func (r *DependencyRepo) DependsOn(parent models.Feature, child models.Feature) bool {
	deps := r.GetDependencies(child)
	_, ok := deps[parent]
	return ok
}

func (r *DependencyRepo) GetDependencies(root models.Feature) models.FeatureSet {
	if _, ok := r.DependencyGraph.DependsOn[root]; !ok {
		return nil
	}

	out := make(models.FeatureSet)
	searchNext := []models.Feature{root}

	for len(searchNext) > 0 {
		discovered := []models.Feature{}
		for _, node := range searchNext {
			dependencyNodes, found := r.DependencyGraph.DependsOn[node]
			if !found {
				continue
			}
			for nextNode := range dependencyNodes {
				if _, ok := out[nextNode]; !ok {
					out[nextNode] = struct{}{}
					discovered = append(discovered, nextNode)
				}
			}
		}
		searchNext = discovered
	}

	return out
}

func (r *DependencyRepo) GetDependents(root models.Feature) models.FeatureSet {
	if _, ok := r.DependencyGraph.Dependents[root]; !ok {
		return nil
	}

	out := make(models.FeatureSet)
	searchNext := []models.Feature{root}
	for len(searchNext) > 0 {
		discovered := []models.Feature{}
		for _, node := range searchNext {
			dependentNodes := r.DependencyGraph.Dependents[node]
			for nextNode := range dependentNodes {
				if _, ok := out[nextNode]; !ok {
					out[nextNode] = struct{}{}
					discovered = append(discovered, nextNode)
				}
			}
		}
		searchNext = discovered
	}

	return out
}
