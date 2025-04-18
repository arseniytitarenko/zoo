package dto

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type AnimalResponse struct {
	ID           uuid.UUID           `json:"id"`
	EnclosureID  uuid.UUID           `json:"enclosure_id"`
	Species      string              `json:"species"`
	Name         string              `json:"name"`
	BirthDate    string              `json:"birth_date"`
	Gender       domain.Gender       `json:"gender"`
	FavoriteFood string              `json:"favorite_food"`
	HealthStatus domain.HealthStatus `json:"health_status"`
}

type NewAnimalRequest struct {
	EnclosureId  uuid.UUID           `json:"enclosure_id" binding:"required"`
	Species      string              `json:"species" binding:"required"`
	Name         string              `json:"name" binding:"required"`
	BirthDate    string              `json:"birth_date" binding:"required"`
	Gender       domain.Gender       `json:"gender" binding:"required"`
	FavoriteFood string              `json:"favorite_food" binding:"required"`
	HealthStatus domain.HealthStatus `json:"health_status" binding:"required"`
}
