package models

type FeatureState string

var (
	ENABLE  FeatureState = "ENABLE"
	DISABLE FeatureState = "DISABLE"
)

type UpdateDependencyRequest struct {
	Parent Feature
	Child  Feature
}

type UpdateFeatureRequest struct {
	Subscription Subscription
	Feature      Feature
	Action       FeatureState
}
