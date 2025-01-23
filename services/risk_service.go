package services

import (
	"errors"
	"github.com/google/uuid"
	"risk_service/models"
	"risk_service/repository"
	"risk_service/utils"
)

type RiskService struct {
	repo repository.RiskRepository
}

func NewRiskService(repo repository.RiskRepository) *RiskService {
	return &RiskService{repo: repo}
}

func (s *RiskService) CreateRisk(risk models.Risk) error {
	if !services.IsValidState(risk.State) {
		return errors.New("invalid state value")
	}
	return s.repo.Create(risk)
}

func (s *RiskService) GetAllRisks() ([]models.Risk, error) {
	return s.repo.GetAll()
}

func (s *RiskService) GetRiskByID(id uuid.UUID) (*models.Risk, error) {
	return s.repo.GetByID(id)
}
