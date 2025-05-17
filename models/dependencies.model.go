package models

type FeatureSet map[Feature]struct{}

type DependencyMap map[Feature]FeatureSet

type DependencyGraph struct {
	FeaturesSet FeatureSet
	DependsOn   DependencyMap // parents
	Dependents  DependencyMap // childs
}
