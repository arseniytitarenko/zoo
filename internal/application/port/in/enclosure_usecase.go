package in

import "zoo/internal/application/dto"

type EnclosureUseCase interface {
	GetAllEnclosures() []dto.EnclosureResponse
	NewEnclosure(req *dto.NewEnclosureRequest) (*dto.EnclosureResponse, error)
	DeleteEnclosure(id string) error
	GetEnclosureByID(id string) (*dto.EnclosureResponse, error)
}
