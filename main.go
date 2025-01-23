package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"risk_service/handlers"
	"risk_service/repository"
	"risk_service/services"
)

func main() {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)
	router := mux.NewRouter()

	router.HandleFunc("/v1/risks", handlers.GetAllRisksHandler(service)).Methods(http.MethodGet)
	router.HandleFunc("/v1/risks", handlers.CreateRiskHandler(service)).Methods(http.MethodPost)
	router.HandleFunc("/v1/risks/{id}", handlers.GetRiskByIDHandler(service)).Methods(http.MethodGet)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
