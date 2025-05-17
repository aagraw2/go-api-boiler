package controller

import (
	"encoding/json"
	"go-api-boiler/models"
	"go-api-boiler/service"
	"net/http"
)

type DependencyController struct {
	serv service.DependencyService
}

func NewDependencyController(s service.DependencyService) *DependencyController {
	return &DependencyController{serv: s}
}

func (uc *DependencyController) UpdateDependency(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UpdateDependencyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := uc.serv.UpdateDependency(req.Parent, req.Child)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
