package repository

import (
	"go-api-boiler/models"
	"testing"
)

func TestEnableFeature(t *testing.T) {

	f1 := models.Feature{
		ID:   "1",
		Name: "Feature1",
	}
	sub1 := models.SubscriptionA{
		ID:       "1",
		Features: []models.Feature{f1},
	}

	f2 := models.Feature{
		ID:   "2",
		Name: "Feature2",
	}

	created, err := sub1.EnableFeature(f2)

	if !created {
		t.Errorf("got false, wanted true")
	}
	if err != nil {
		t.Errorf("got err %s", err.Error())
	}
}

func TestDisableFeature(t *testing.T) {

	f1 := models.Feature{
		ID:   "1",
		Name: "Feature1",
	}
	sub1 := models.SubscriptionA{
		ID:       "1",
		Features: []models.Feature{f1},
	}

	created, err := sub1.DisableFeature(f1)

	if !created {
		t.Errorf("got false, wanted true")
	}
	if err != nil {
		t.Errorf("got err %s", err.Error())
	}
}
