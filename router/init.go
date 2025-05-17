package router

import (
	"go-api-boiler/controller"
	"go-api-boiler/repository"
	"go-api-boiler/service"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	clientRepo := repository.NewClientRepository()
	dependencyRepository := repository.NewDependencyRepository()
	featureRepository := repository.NewFeatureRepository(clientRepo)

	featureService := service.NewFeatureService(featureRepository, clientRepo, dependencyRepository)
	featureController := controller.NewFeatureController(featureService)

	mux.HandleFunc("/api/feature/update", featureController.UpdateFeature)

	dependencyService := service.NewDependencyService(dependencyRepository)
	dependencyController := controller.NewDependencyController(dependencyService)

	mux.HandleFunc("/api/dependency/add", dependencyController.UpdateDependency)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, pattern := mux.Handler(r)
		if pattern == "" {
			// Custom 404 Error
			http.Error(w, "Route not found", http.StatusNotFound)
			return
		}
		handler.ServeHTTP(w, r)
	})

}
