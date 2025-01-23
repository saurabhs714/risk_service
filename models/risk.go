package models

import "github.com/google/uuid"

type Risk struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
}

// NewUUID generates a new UUID and returns it.
func NewUUID() uuid.UUID {
	return uuid.New()
}
