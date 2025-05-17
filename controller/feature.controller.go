package controller

import (
	"encoding/json"
	"go-api-boiler/models"
	"go-api-boiler/service"
	"net/http"
)

type FeatureController struct {
	serv service.FeatureService
}

func NewFeatureController(s service.FeatureService) *FeatureController {
	return &FeatureController{serv: s}
}

func (uc *FeatureController) UpdateFeature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UpdateFeatureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	done, err := uc.serv.UpdateFeature(req.Subscription, req.Feature, req.Action)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(done)
}
