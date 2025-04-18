package repository

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type IAnimalRepository interface {
	Save(animal *domain.Animal)
	Delete(id uuid.UUID)
	GetByID(id uuid.UUID) (*domain.Animal, bool)
	GetAll() []domain.Animal
}
