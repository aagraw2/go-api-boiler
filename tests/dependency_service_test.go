package repository

import (
	"go-api-boiler/models"
	"go-api-boiler/repository"
	"testing"
)

func TestAddRelationHappyFlow(t *testing.T) {

	DependencyRepo := repository.NewDependencyRepository()

	f1 := models.Feature{
		ID:   "1",
		Name: "Feature1",
	}
	f2 := models.Feature{
		ID:   "2",
		Name: "Feature2",
	}

	created, err := DependencyRepo.AddDependency(f1, f2)
	if !created {
		t.Errorf("got false, wanted true")
	}
	if err != nil {
		t.Errorf("got err %s", err.Error())
	}
}

func TestAddRelationSelfDependency(t *testing.T) {

	DependencyRepo := repository.NewDependencyRepository()

	f1 := models.Feature{
		ID:   "1",
		Name: "Feature1",
	}

	_, err := DependencyRepo.AddDependency(f1, f1)

	if err == nil || err.Error() != "cannot depend on self" {
		t.Errorf("got err %s", err.Error())
	}
}

func TestRelationCircularDependency(t *testing.T) {
	DependencyRepo := repository.NewDependencyRepository()

	f1 := models.Feature{
		ID:   "1",
		Name: "Feature1",
	}
	f2 := models.Feature{
		ID:   "2",
		Name: "Feature2",
	}
	f3 := models.Feature{
		ID:   "3",
		Name: "Feature3",
	}
	f4 := models.Feature{
		ID:   "4",
		Name: "Feature3",
	}

	_, _ = DependencyRepo.AddDependency(f1, f2)
	_, _ = DependencyRepo.AddDependency(f2, f3)
	_, _ = DependencyRepo.AddDependency(f2, f4)
	_, err := DependencyRepo.AddDependency(f4, f1)

	if err == nil || err.Error() != "circular dependencies not allowed" {
		t.Errorf("got err %s", err.Error())
	}
}
