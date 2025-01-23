package repository

import (
	"github.com/google/uuid"
	"risk_service/models"
	"sync"
)

type RiskRepository interface {
	Create(risk models.Risk) error
	GetAll() ([]models.Risk, error)
	GetByID(id uuid.UUID) (*models.Risk, error)
}

type InMemoryRiskRepository struct {
	data map[uuid.UUID]models.Risk
	mu   sync.RWMutex
}

func NewInMemoryRiskRepository() *InMemoryRiskRepository {
	return &InMemoryRiskRepository{
		data: make(map[uuid.UUID]models.Risk),
	}
}

func (r *InMemoryRiskRepository) Create(risk models.Risk) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[risk.ID] = risk
	return nil
}

func (r *InMemoryRiskRepository) GetAll() ([]models.Risk, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	risks := make([]models.Risk, 0, len(r.data))
	for _, risk := range r.data {
		risks = append(risks, risk)
	}
	return risks, nil
}

func (r *InMemoryRiskRepository) GetByID(id uuid.UUID) (*models.Risk, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	risk, exists := r.data[id]
	if !exists {
		return nil, nil
	}
	return &risk, nil
}
