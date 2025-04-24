package router

import (
	"go-api-boiler/controller"
	"go-api-boiler/repository"
	"go-api-boiler/service"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	mux.HandleFunc("/api/user/get", userController.GetUsers)
	mux.HandleFunc("/api/user/create", userController.CreateUsers)
	return mux
}
