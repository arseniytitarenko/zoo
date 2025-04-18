package in

import (
	"zoo/internal/application/dto"
)

type AnimalUseCase interface {
	GetAllAnimals() []dto.AnimalResponse
	NewAnimal(req *dto.NewAnimalRequest) (*dto.AnimalResponse, error)
	DeleteAnimal(id string) error
	GetAnimalByID(id string) (*dto.AnimalResponse, error)
}
