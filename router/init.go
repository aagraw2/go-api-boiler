package router

import (
	"go-api-boiler/controller"
	"go-api-boiler/repository"
	"go-api-boiler/service"
	"go-api-boiler/utils"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	sharedDB := utils.GetDBInstance()
	userRepository := repository.NewUserRepository(sharedDB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	mux.HandleFunc("/api/user/get", userController.GetUsers)
	mux.HandleFunc("/api/user/create", userController.CreateUsers)

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
