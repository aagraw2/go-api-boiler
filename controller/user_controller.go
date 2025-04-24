package controller

import (
	"encoding/json"
	"go-api-boiler/models"
	"go-api-boiler/service"
	"net/http"
)

type UserController struct {
	serv service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{serv: s}
}

func (uc *UserController) CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created := uc.serv.CreateUser(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	users := uc.serv.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
