package out

import (
	"github.com/google/uuid"
	"zoo/internal/domain"
)

type AnimalRepository interface {
	Save(animal *domain.Animal)
	Delete(id uuid.UUID)
	GetByID(id uuid.UUID) (*domain.Animal, bool)
	GetAll() []domain.Animal
}
