package dto

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type EnclosureResponse struct {
	ID              uuid.UUID            `json:"id"`
	Type            domain.EnclosureType `json:"type"`
	Length          float64              `json:"length"`
	Width           float64              `json:"width"`
	Height          float64              `json:"height"`
	CurrAnimalCount uint8                `json:"curr_animal_count"`
	AnimalCapacity  uint8                `json:"animal_capacity"`
}

type NewEnclosureRequest struct {
	Type           domain.EnclosureType `json:"type"`
	Length         float64              `json:"length"`
	Width          float64              `json:"width"`
	Height         float64              `json:"height"`
	AnimalCapacity uint8                `json:"animal_capacity"`
}
