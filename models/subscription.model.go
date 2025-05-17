package models

type Subscription interface {
	GetEnabledFeatures() map[Feature]bool
	EnableFeature(feature Feature) (bool, error)
	DisableFeature(feature Feature) (bool, error)
}

type SubscriptionA struct {
	ID       string
	Plan     string
	Price    int
	Features []Feature
}

func (a *SubscriptionA) GetEnabledFeatures() map[Feature]bool {
	featureMap := make(map[Feature]bool)
	for _, f := range a.Features {
		featureMap[f] = true
	}
	return featureMap
}

func (a *SubscriptionA) EnableFeature(feature Feature) (bool, error) {
	a.Features = append(a.Features, feature)
	return true, nil
}

func (a *SubscriptionA) DisableFeature(feature Feature) (bool, error) {
	for i, f := range a.Features {
		if feature == f {
			a.Features = append(a.Features[:i], a.Features[i+1:]...)
		}
	}
	return true, nil
}
