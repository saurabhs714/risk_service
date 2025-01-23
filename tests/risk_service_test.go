package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"risk_service/handlers"
	"risk_service/models"
	"risk_service/repository"
	"risk_service/services"
	"testing"
)

func TestCreateRiskHandler(t *testing.T) {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)

	t.Run("Create valid risk", func(t *testing.T) {
		risk := map[string]string{
			"title":       "Valid Risk",
			"description": "A valid risk description",
			"state":       "open",
		}
		body, _ := json.Marshal(risk)
		req := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handlers.CreateRiskHandler(service)(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Valid Risk")
	})

	t.Run("Create risk with invalid state", func(t *testing.T) {
		risk := map[string]string{
			"title":       "Invalid Risk",
			"description": "Risk with invalid state",
			"state":       "invalid_state",
		}
		body, _ := json.Marshal(risk)
		req := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		handlers.CreateRiskHandler(service)(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid state value")
	})
}

func TestGetAllRisksHandler(t *testing.T) {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)

	// Seed some test data
	service.CreateRisk(models.Risk{
		ID:          models.NewUUID(),
		Title:       "Risk 1",
		Description: "Test Risk 1",
		State:       "open",
	})
	service.CreateRisk(models.Risk{
		ID:          models.NewUUID(),
		Title:       "Risk 2",
		Description: "Test Risk 2",
		State:       "closed",
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/risks", nil)
	w := httptest.NewRecorder()

	handlers.GetAllRisksHandler(service)(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Risk 1")
	assert.Contains(t, w.Body.String(), "Risk 2")
}

func TestGetRiskByIDHandler(t *testing.T) {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)
	router := mux.NewRouter()
	router.HandleFunc("/v1/risks/{id}", handlers.GetRiskByIDHandler(service))

	// Seed a test risk
	risk := models.Risk{
		ID:          models.NewUUID(),
		Title:       "Risk To Fetch",
		Description: "Risk for testing Get by ID",
		State:       "investigating",
	}
	service.CreateRisk(risk)

	t.Run("Get existing risk by ID", func(t *testing.T) {
		url := "/v1/risks/" + risk.ID.String()
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		// Serve the request through the router
		router.ServeHTTP(w, req)

		// Debug: Log response and mux vars
		t.Logf("Handler Response Body: %s", w.Body.String())
		t.Logf("Handler Status Code: %d", w.Code)

		// Assert response
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), risk.Title)
	})

	t.Run("Get non-existing risk by ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/risks/non-existent-id", nil)
		w := httptest.NewRecorder()

		handlers.GetRiskByIDHandler(service)(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid ID format")
	})
}

func TestCreateRiskHandler_InvalidJSON(t *testing.T) {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)

	req := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	handlers.CreateRiskHandler(service)(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Bad request")
}

func TestHandlersIntegration(t *testing.T) {
	repo := repository.NewInMemoryRiskRepository()
	service := services.NewRiskService(repo)

	t.Run("End-to-end test: Create and Fetch Risk", func(t *testing.T) {
		// Create risk
		risk := map[string]string{
			"title":       "E2E Risk",
			"description": "E2E Test Risk",
			"state":       "accepted",
		}
		body, _ := json.Marshal(risk)
		createReq := httptest.NewRequest(http.MethodPost, "/v1/risks", bytes.NewBuffer(body))
		createResp := httptest.NewRecorder()
		handlers.CreateRiskHandler(service)(createResp, createReq)

		assert.Equal(t, http.StatusCreated, createResp.Code)

		// Fetch all risks
		getReq := httptest.NewRequest(http.MethodGet, "/v1/risks", nil)
		getResp := httptest.NewRecorder()
		handlers.GetAllRisksHandler(service)(getResp, getReq)

		assert.Equal(t, http.StatusOK, getResp.Code)
		assert.Contains(t, getResp.Body.String(), "E2E Risk")
	})
}
