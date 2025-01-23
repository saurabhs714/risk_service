package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"risk_service/models"
	"risk_service/services"
)

func GetAllRisksHandler(service *services.RiskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		risks, err := service.GetAllRisks()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(risks)
	}
}

func CreateRiskHandler(service *services.RiskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var risk models.Risk
		if err := json.NewDecoder(r.Body).Decode(&risk); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		risk.ID = uuid.New()
		if err := service.CreateRisk(risk); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(risk)
	}
}

func GetRiskByIDHandler(service *services.RiskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		risk, err := service.GetRiskByID(id)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if risk == nil {
			http.Error(w, "Risk not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(risk)
	}
}
